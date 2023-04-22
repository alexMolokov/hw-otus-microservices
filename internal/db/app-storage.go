package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	"github.com/jmoiron/sqlx"
)

var ErrUserNameExists = errors.New("user name exists")

type AppStorage struct {
	db *sqlx.DB
}

func (s *AppStorage) UserCreate(ctx context.Context, request model.UserCreateRequest) (int64, error) {
	ctx, cancelCtx := s.getTimeoutContext(ctx)
	defer cancelCtx()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	row := tx.QueryRowContext(ctx,
		`SELECT user_id FROM main.user WHERE user_name = $1 `, request.UserName)

	if err = row.Err(); err != nil {
		return 0, err
	}

	var id int64
	err = row.Scan(&id)
	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		return 0, ErrUserNameExists
	}

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO main.user (user_name, first_name, last_name, email, phone)
			   VALUES ($1,$2,$3,$4,$5) RETURNING user_id`,
	)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	row = stmt.QueryRowContext(ctx, request.UserName, request.FirstName, request.LastName,
		request.Email, request.Phone,
	)
	if row.Err() != nil {
		return 0, row.Err()
	}
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *AppStorage) UserUpdate(ctx context.Context, request model.UserUpdateRequest) error {
	ctx, cancelCtx := s.getTimeoutContext(ctx)
	defer cancelCtx()

	_, err := s.db.ExecContext(ctx,
		`UPDATE main.user 
				SET first_name = $2, last_name = $3, email = $4, phone = $5
                WHERE user_id = $1`, request.UserID, request.FirstName, request.LastName,
		request.Email, request.Phone,
	)

	return err
}

func (s *AppStorage) UserDelete(ctx context.Context, id int64) error {
	ctx, cancelCtx := s.getTimeoutContext(ctx)
	defer cancelCtx()

	_, err := s.db.ExecContext(ctx, `DELETE FROM main.user WHERE user_id = $1 `, id)
	return err
}

func (s *AppStorage) UserGet(ctx context.Context, id int64) (*model.User, error) {
	ctx, cancelCtx := s.getTimeoutContext(ctx)
	defer cancelCtx()

	row := s.db.QueryRowxContext(ctx,
		`SELECT user_id, user_name, first_name, last_name, email, phone 
			FROM main.user 
			WHERE user_id = $1 `, id)

	if err := row.Err(); err != nil {
		return nil, err
	}

	u := model.User{}
	if err := row.Scan(&u.UserID, &u.UserName, &u.FirstName, &u.LastName, &u.Email, &u.Phone); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *AppStorage) getTimeoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 3*time.Second)
}

func NewAppStorage(connection *Connection) *AppStorage {
	return &AppStorage{
		db: connection.GetDB(),
	}
}
