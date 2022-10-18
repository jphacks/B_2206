package domain

import (
	"time"
)

type Range struct {
	//data
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//belongs to
	MaxValueId int `json:"max_value_id" gorm:"foreign_key:ID"`
	MinValueId int `json:"min_value_id" gorm:"foreign_key:ID"`
	//has many
	DetailsRanges []DetailsRange `json:"details_ranges"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Ranges []Range
