package domain

import (
	"time"
)

type DetailsTag struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//belongs to
	DetailId         int `json:"detail_id" gorm:"foreign_key:ID"`
	TagId            int `json:"tag_id" gorm:"foreign_key:ID"`
	ClassificationId int `json:"classification_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type DetailsTags []DetailsTag
