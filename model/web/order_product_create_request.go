package web

import "time"

type OrderProductCreateRequest struct {
	OrderId   int       `validate:"required,min=1,max=100" json:"order_id"`
	ProductId int       `Validate:"required,min=1,max=100" json:"product_id"`
	Qty       int       `Validate:"required,min=1,max=100" json:"qty"`
	Price     int       `Validate:"required,min=1,max=100" json:"price"`
	Amount    int       `Validate:"required,min=1,max=100" json:"amount"`
	CreatedAt time.Time `Validate:"required,min=1,max=100" json:"created_at"`
	UpdatedAt time.Time `Validate:"required,min=1,max=100" json:"updated_at"`
}
