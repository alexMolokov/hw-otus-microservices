-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS main;

CREATE TABLE IF NOT EXISTS main.user
(
    user_id         BIGSERIAL PRIMARY KEY,
    user_name       VARCHAR(255) UNIQUE NOT NULL,
    first_name      VARCHAR(255),
    last_name       VARCHAR(255),
    email           VARCHAR(255),
    phone           VARCHAR(20),
    created_at      TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now()
);

COMMENT ON TABLE main.user IS 'Пользователи системы';
COMMENT ON COLUMN main.user.user_name IS 'Имя пользователя (логин)';
COMMENT ON COLUMN main.user.first_name IS 'Фамилия';
COMMENT ON COLUMN main.user.last_name IS 'Имя';
COMMENT ON COLUMN main.user.email IS 'Email';
COMMENT ON COLUMN main.user.phone IS 'Телефон';
COMMENT ON COLUMN main.user.created_at IS 'Дата создания записи';
COMMENT ON COLUMN main.user.updated_at IS 'Дата изменения записи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS main.user;
-- +goose StatementEnd
