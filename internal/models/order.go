package models

import "time"

type Order struct {
	Id         int64     `json:"id"`
	CourierId  int64     `json:"courier_id"`
	RecieverId int64     `json:"reciver_id"`
	Status     string    `json:"status"`
	OrderedAt  time.Time `json:"ordered_at"`
}

type OrderCreateRequest struct {
	RecieverId int64   `json:"-"`
	LocationX  float64 `json:"location_x"`
	LocationY  float64 `json:"location_y"`
	Address    string  `json:"address"`

	Products []OrderProductCreateRequest `json:"products"`
}

type OrderProduct struct {
	Id        int64 `json:"id"`
	OrderId   int64 `json:"order_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type OrderProductCreateRequest struct {
	ProductId  int64                `json:"product_id"`
	Attributes []AttributeAndOption `json:"attributes"`
	Quantity   int                  `json:"quantity"`
}
