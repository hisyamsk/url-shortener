package models

type UserModel struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=5,max=12"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserUpdateRequest struct {
	Username string `json:"username" validate:"omitempty,min=5,max=12"`
	Password string `json:"password" validate:"omitempty,min=6,max=50"`
}
