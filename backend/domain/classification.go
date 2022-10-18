package domain

import (
	"time"
)

type Classification struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	DetailsRanges []DetailsRange
	DetailsTags []DetailsTag
	DetailsValues []DetailsValue
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Classifications []Classification