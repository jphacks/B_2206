package domain

import (
	"time"
)

type Matching struct {
	//data
	ID            int    `json:"id" gorm:"primary_key"`
	Status        string `json:"status"`
	SellerMessage string `json:"seller_message"`
	//belongs to
	UserId    int `json:"user_id" gorm:"foreign_key:ID"`
	RequestId int `json:"request_id" gorm:"foreign_key:ID"`
	BuyerId   int `json:"buyer_user_id" gorm:"foreign_key:ID"`
	SellerId  int `json:"seller_user_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Matchings []Matching
