package models

type User struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	RoleId   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email" default:"khdiyz.12@gmail.com"`
	Password string `json:"password" validate:"required" default:"Secret@12"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserCreateRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
	RoleId   int64  `json:"role_id"`
}
