package domain

import (
	"time"
)

type Value struct {
	//data
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Value string `json:"value"`
	//has many
	Ranges []Range `json:"ranges"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Values []Value
