package model

type UserCreateRequest struct {
	UserName string `json:"username" valid:"required~Поле username обязательно для заполнения" example:"alex.molokov"`
	userCommon
}

type UserUpdateRequest struct {
	UserID int64
	userCommon
}

type User struct {
	UserID int64 `json:"id"`
	UserCreateRequest
}

type userCommon struct {
	FirstName *string `json:"firstName" example:"Молоков"`
	LastName  *string `json:"lastName" example:"Алексей"`
	Email     *string `json:"email" example:"alex.molokov@yandex.ru"`
	Phone     *string `json:"phone" example:"+79035431754"`
}
