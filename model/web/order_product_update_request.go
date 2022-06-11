package web

type OrderProductUpdateRequest struct {
	Id        int `validate:"required"`
	OrderId   int `validate:"required,max=200,min=1" json:"order_id"`
	ProductId int `validate:"required,max=200,min=1" json:"product_id"`
	Qty       int `validate:"required,max=200,min=1" json:"qty"`
	Price     int `validate:"required,max=200,min=1" json:"price"`
	Amount    int `validate:"required,max=200,min=1" json:"amount"`
}
