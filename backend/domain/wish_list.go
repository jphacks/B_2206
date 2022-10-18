package domain

import (
	"time"
)

type WishList struct {
	//data
	ID				int				`json:"id" gorm:"primary_key"`
	Name 	 		string		`json:"name"`
	//belongs to
	UserId 		int				`json:"user_id" gorm:"foreign_key:ID"`
	RequestId 	int				`json:"request_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt	time.Time	`json:"updated_at"`
	CreatedAt	time.Time	`json:"created_at"`
}

type WishLists []WatchList