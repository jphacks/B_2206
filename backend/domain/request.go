package domain

import (
	"time"
)

type Request struct {
	ID int `json:"id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null"`
	IsPurchase bool `json:"is_purchase" gorm:"not null"`
	User User
	Detail Detail
	WishList []WishList
	Matchings []Matching
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Requests []Request