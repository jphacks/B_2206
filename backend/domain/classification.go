package domain

import (
	"time"
)

type Classification struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//has many
	DetailsRanges []DetailsRange
	DetailsTags []DetailsTag
	DetailsValues []DetailsValue
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Classifications []Classification