package domain

import (
	"time"
)

type DetailsValue struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//belongs to
	DetailId int `json:"detail_id" gorm:"foreign_key:ID"`
	ValueId int `json:"value_id" gorm:"foreign_key:ID"`
	ClassificationId int `json:"classification_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type DetailsValues []DetailsValue