package domain

import (
	"time"
)

type Tag struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	DetailsTags []DetailsTag
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Tags []Tag