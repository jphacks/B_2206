package domain

import (
	"time"
)

type Area struct {
	//data
	ID            int    `json:"id" gorm:"primary_key"`
	PostCode      string `json:"post_code" gorm:"not null"`
	Prefecture    string `json:"prefecture" gorm:"not null"`
	City          string `json:"city" gorm:"not null"`
	AddressNumber string `json:"address_number"`
	BuildingName  string `json:"building_name"`
	//belongs to
	DetailId int `json:"detail_id" gorm:"foreign_key:ID"`
	// //has one
	// Detail Detail `json:"detail"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Areas []Area
