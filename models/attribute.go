package models

type Attribute struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type AttributeCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type AttributeUpdateRequest struct {
	Id   int64  `json:"-"`
	Name string `json:"name" validate:"required"`
}
