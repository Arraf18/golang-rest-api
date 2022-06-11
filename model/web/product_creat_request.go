package web

import "time"

type ProductCreateRequest struct {
	Name       string    `validate:"required,min=1,max=100" json:"name"`
	Price      int       `validate:"required,min=1,max=100" json:"price"`
	CategoryId int       `validate:"required,min=1,max=100" json:"category_id"`
	CreatedAt  time.Time `Validate:"required,min=1,max=100" json:"created_at"`
	UpdatedAt  time.Time `Validate:"required,min=1,max=100" json:"updated_at"`
}
