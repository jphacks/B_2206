package domain

import (
	"time"
)

type Request struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null"`
	IsPurchase bool `json:"is_purchase" gorm:"not null"`
	//belongs to
	UserId int `json:"user_id" gorm:"foreign_key:ID"`
	DetailId int `json:"detail_id" gorm:"foreign_key:ID"`
	//has many
	WishLists []WishList
	Matchings []Matching
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Requests []Request