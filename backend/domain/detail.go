package domain

import (
	"time"
)

type Detail struct {
	//data
	ID int `json:"id" gorm:"primary_key"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// //has one
// Request    Request   `json:"request"`
// WatchiList WatchList `json:"watch_list"`
// //has many
// DetailsRanges []DetailsRange `json:"details_ranges"`
// DetailsTags   []DetailsTag   `json:"details_tags"`
// DetailsValues []DetailsValue `json:"details_values"`

type DetailAll struct {
	Areas         []Area
	DetailsLimits DetailsLimits
	DetailsValues DetailsValues
	DetailsTags   DetailsTags
}

type Details []Detail
