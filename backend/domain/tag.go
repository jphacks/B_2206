package domain

import (
	"time"
)

type Tag struct {
	//data
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//has many
	DetailsTags []DetailsTag `json:"details_tags"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Tags []Tag
