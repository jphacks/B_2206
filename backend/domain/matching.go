package domain

import (
	"time"
)

type Matching struct {
	ID int `json:"id" gorm:"primary_key"`
	Request Request
	BuyerID User
	SellerID User
	Status string `json:"status"`
	SellerMessage string `json:"seller_message"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Matchings []Matching