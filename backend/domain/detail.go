package domain

import (
	"time"
)

type Detail struct {
	ID int `json:"id" gorm:"primary_key"`
	Area Area
	DetailsRange []DetailsRange
	DetailsTag []DetailsTag
	DetailsValue []DetailsValue
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Details []Detail