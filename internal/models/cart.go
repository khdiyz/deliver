package models

type Cart struct {
	Id     int64
	UserId int64
}

type CartProductCreateRequest struct {
	CartId     int64
	ProductId  int64
	Quantity   int
	Attributes []AttributeAndOption
}
