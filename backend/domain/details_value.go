package domain

import (
	"time"
)

type DetailsValue struct {
	ID int `json:"id" gorm:"primary_key"`
	Detail Detail
	Value Value
	Classification Classification
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type DetailsValues []DetailsValue