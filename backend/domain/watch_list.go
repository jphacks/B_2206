package domain

import (
	"time"
)

type WatchList struct {
	//data
	ID			int			`json:"id" gorm:"primary_key"`
	Description string		`json:"description"`
	IsPurchase	bool		`json:"is_purchase"`
	//belongs to
	UserId		int			`json:"user_id" gorm:"foreign_key:ID"`
	DetailId	int			`json:"detail_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt	time.Time	`json:"updated_at"`
	CreatedAt	time.Time	`json:"created_at"`
}

type WatchLists []WatchList