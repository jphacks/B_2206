package domain

import (
	"time"
)

type DetailsRange struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//belongs to
	DetailId         int    `json:"detail_id" gorm:"foreign_key:ID"`
	RangeId          int    `json:"range_id" gorm:"foreign_key:ID"`
	ClassificationId string `json:"classification_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type DetailsRanges []DetailsRange
