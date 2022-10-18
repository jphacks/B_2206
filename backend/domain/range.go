package domain

import (
	"time"
)

type Range struct {
	//data
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//belongs to
	ValueId int `json:"value_id" gorm:"foreign_key:ID"`
	// MinValueId int `json:"value_id" gorm:"foreign_key:ID"`
	//has many
	// Values        []Value        `json:"values"`
	DetailsRanges []DetailsRange `json:"details_ranges"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Ranges []Range
