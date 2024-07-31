package models

type Product struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	CategoryId  int64  `json:"category_id"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Price       int    `json:"price"`
}

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	CategoryId  int64  `json:"category_id"`
	Description string `json:"description"`
	Photo       string `json:"photo" validate:"required"`
	Price       int    `json:"price"`
}

type ProductUpdateRequest struct {
	Id          int64  `json:"-"`
	Name        string `json:"name" validate:"required"`
	CategoryId  int64  `json:"category_id"`
	Description string `json:"description"`
	Photo       string `json:"photo" validate:"required"`
	Price       int    `json:"price"`
}
