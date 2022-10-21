package domain

import (
	"time"
)

type PersonalInfo struct {
	//data
	ID             int    `json:"id" gorm:"primary_key"`
	FamilyName     string `json:"family_name"`
	FirstName      string `json:"first_name"`
	FamilyNameKana string `json:"family_name_kana"`
	FirstNameKana  string `json:"first_name_kana"`
	Birthday       string `json:"birthday"`
	PhoneNumber    string `json:"phone_number"`
	//belongs_to
	UserID int `json:"user_id"`
	//time stamp
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type PersonalInfos []PersonalInfo
