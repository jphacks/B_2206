package domain

import (
	"time"
)

type Detail struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//belongs to
	AreaId int `json:"area_id" gorm:"foreign_key:ID"`
	//has one
	Request Request
	WatchiList WatchList
	WishList WishList
	//has many
	DetailsRanges []DetailsRange
	DetailsTags []DetailsTag
	DetailsValues []DetailsValue
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Details []Detail