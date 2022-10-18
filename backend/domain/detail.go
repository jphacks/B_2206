package domain

import (
	"time"
)

type Detail struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//belongs to
	AreaId int `json:"area_id" gorm:"foreign_key:ID"`
	//has one
	Request    Request   `json:"request"`
	WatchiList WatchList `json:"watch_list"`
	//has many
	DetailsRanges []DetailsRange `json:"details_ranges"`
	DetailsTags   []DetailsTag   `json:"details_tags"`
	DetailsValues []DetailsValue `json:"details_values"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Details []Detail
