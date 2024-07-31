package models

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id   int64  `json:"-"`
	Name string `json:"name" validate:"required"`
}
