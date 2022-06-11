package web

import "time"

type OrdersCreateRequest struct {
	OrderDate   time.Time `validate:"required,min=1,max=100" json:"order_date"`
	CustomerId  int       `Validate:"required,min=1,max=100" json:"customer_id"`
	TotalAmount int       `Validate:"required,min=1,max=100" json:"total_amount"`
	CreatedAt   time.Time `Validate:"required,min=1,max=100" json:"created_at"`
	UpdatedAt   time.Time `Validate:"required,min=1,max=100" json:"updated_at"`
}
