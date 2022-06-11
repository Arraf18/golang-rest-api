package web

import "time"

type OrdersResponse struct {
	Id          int       `json:"id"`
	OrderDate   time.Time `json:"order_date"`
	CustomerId  int       `json:"customer_id"`
	TotalAmount int       `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
