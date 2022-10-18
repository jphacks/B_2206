package domain

import (
	"time"
)

type DetailsTag struct {
	ID int `json:"id" gorm:"primary_key"`
	Detail Detail
	Tag Tag
	Classification Classification
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type DetailsTags []DetailsTag