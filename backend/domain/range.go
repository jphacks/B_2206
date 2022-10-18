package domain

import (
	"time"
)

type Range struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	MaxValue Value `json:"max_value_id"`
	MinValue Value `json:"min_value_id"`
	Classification Classification
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Ranges []Range