package web

import "time"

type CustomerResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
