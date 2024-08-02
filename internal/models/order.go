package models

import "time"

type Order struct {
	Id           int64     `json:"id"`
	CourierId    int64     `json:"courier_id"`
	RecieverId   int64     `json:"reciver_id"`
	RecieverName string    `json:"reciever_name,omitempty"`
	LocationX    float64   `json:"location_x"`
	LocationY    float64   `json:"location_y"`
	Address      string    `json:"address"`
	Status       string    `json:"status"`
	OrderedAt    time.Time `json:"ordered_at"`

	Products []OrderProduct `json:"products,omitempty"`
}

type OrderCreateRequest struct {
	RecieverId int64   `json:"-"`
	LocationX  float64 `json:"location_x"`
	LocationY  float64 `json:"location_y"`
	Address    string  `json:"address"`

	Products []OrderProductCreateRequest `json:"products"`
}

type OrderUpdateRequest struct {
	Id         int64   `json:"-"`
	RecieverId int64   `json:"-"`
	CourierId  int64   `json:"courier_id"`
	LocationX  float64 `json:"location_x"`
	LocationY  float64 `json:"location_y"`
	Address    string  `json:"address"`
	Status     string  `json:"status"`
}

type OrderProduct struct {
	Id          int64  `json:"id,omitempty"`
	OrderId     int64  `json:"order_id,omitempty"`
	ProductId   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`

	Attributes interface{} `json:"attributes"`
}

type OrderProductCreateRequest struct {
	ProductId  int64                `json:"product_id"`
	Attributes []AttributeAndOption `json:"attributes"`
	Quantity   int                  `json:"quantity"`
}

type OrderCourierRequest struct {
	CourierID int64  `json:"courier_id"`
	Status    string `json:"status"`
}
