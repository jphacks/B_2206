package domain

import (
	"time"
)

type PersonalInfo struct {
	//data
	ID          int    `json:"id" gorm:"primary_key"`
	FamilyName  string `json:"family_name" gorm:"not null"`
	FirstName   string `json:"first_name" gorm:"not null"`
	Birthday    string `json:"birthday" gorm:"not null"`
	PhoneNumber int    `json:"phone_number" gorm:"not null"`
	//has one
	User User `json:"user"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type PersonalInfos []PersonalInfo
