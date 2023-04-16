package model

type StatusResponse struct {
	Status string `json:"status" example:"OK"`
}

type UserCreateResponse struct {
	Error   bool   `json:"error" example:"false"`
	ID      int64  `json:"id"  example:"1"`
	Message string `json:"message"  example:"OK"`
}

func NewUserCreateResponse(id int64) *UserCreateResponse {
	return &UserCreateResponse{
		Error:   false,
		Message: "User has been created",
		ID:      id,
	}
}
