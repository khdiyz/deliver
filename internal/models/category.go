package models

type Category struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products,omitempty"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id   int64  `json:"-"`
	Name string `json:"name" validate:"required"`
}
