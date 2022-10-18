package domain

import (
	"time"
)

type WishList struct {
	//data
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//belongs to
	RequestId  int `json:"request_id" gorm:"foreign_key:ID"`
	WishListId int `json:"wish_list_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type WishLists []WishList
