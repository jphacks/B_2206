package domain

import (
	"time"
)

type WatchList struct {
	ID			int			`json:"id" gorm:"primary_key"`
	User		User
	Detail		Detail
	UpdatedAt	time.Time	`json:"updated_at"`
	CreatedAt	time.Time	`json:"created_at"`
}

type WatchLists []WatchList