package domain

import (
	"time"
)

type CompanyInfo struct {
	//data
	ID            int    `json:"id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"not null"`
	PhoneNumber   string `json:"phone_number" gorm:"not null"`
	PostCode      string `json:"post_code" gorm:"not null"`
	AddressNumber string `json:"address_number" gorm:"not null"`
	BuildingName  string `json:"building_name"`
	Website       string `json:"website"`
	// //has one
	// User User `json:"user"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CompanyInfos []CompanyInfo
