package web

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Name       string `validate:"required,max=200,min=1" json:"name"`
	Price      int    `validate:"required,max=200,min=1" json:"price"`
	CategoryId int    `validate:"required,max=200,min=1" json:"category_id"`
}
