package domain

import (
	"time"
)

type Detail struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	// //has one
	// Request    Request   `json:"request"`
	// WatchiList WatchList `json:"watch_list"`
	// //has many
	// DetailsRanges []DetailsRange `json:"details_ranges"`
	// DetailsTags   []DetailsTag   `json:"details_tags"`
	// DetailsValues []DetailsValue `json:"details_values"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Details []Detail

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

type DetailsValue struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//belongs to
	DetailId         int `json:"detail_id" gorm:"foreign_key:ID"`
	ValueId          int `json:"value_id" gorm:"foreign_key:ID"`
	ClassificationId int `json:"classification_id" gorm:"foreign_key:ID"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type DetailsValues []DetailsValue
