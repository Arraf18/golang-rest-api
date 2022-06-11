package web

import "time"

type CustomerCreateRequest struct {
	Name        string    `validate:"required,min=1,max=100" json:"name"`
	Address     string    `Validate:"required,min=1,max=100" json:"address"`
	Email       string    `Validate:"required,min=1,max=100" json:"email"`
	PhoneNumber string    `Validate:"required,min=1,max=100" json:"phone_number"`
	CreatedAt   time.Time `Validate:"required,min=1,max=100" json:"created_at"`
	UpdatedAt   time.Time `Validate:"required,min=1,max=100" json:"updated_at"`
}
