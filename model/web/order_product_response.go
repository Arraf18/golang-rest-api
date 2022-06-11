package web

import "time"

type OrderProductResponse struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"order_id"`
	ProductId int       `json:"product_id"`
	Qty       int       `json:"qty"`
	Price     int       `json:"price"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
