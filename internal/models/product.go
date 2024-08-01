package models

type Product struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	CategoryId  int64       `json:"category_id,omitempty"`
	Description string      `json:"description"`
	Photo       string      `json:"photo"`
	Price       int         `json:"price"`
	Attributes  []Attribute `json:"attributes,omitempty"`
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

type AddAttributeToProduct struct {
	ProductId   int64 `json:"product_id"`
	AttributeId int64 `json:"attribute_id"`
}

type ProductAttribute struct {
	Id          int64 `json:"id"`
	ProductId   int64 `json:"product_id"`
	AttributeId int64 `json:"attribute_id"`
}

type AttributeAndOption struct {
	AttributeId int64 `json:"attribute_id"`
	OptionId    int64 `json:"option_id"`
}

type AddToCart struct {
	Quantity   int                  `json:"quantity"`
	Attributes []AttributeAndOption `json:"attributes"`
}
