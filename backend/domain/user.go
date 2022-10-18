package domain

import (
	"time"
)

type User struct {
	//data
	ID						int    			`json:"id" gorm:"primary_key"`
	Name					string			`json:"name" gorm:"not null"`
	Email					string			`json:"email" gorm:"not null"`
	Password			string			`json:"password" gorm:"not null"`
	//belongs to
	PersonalInfoId	int				`json:"personal_info_id" gorm:"foreign_key:ID"`
	CompanyInfoId	int					`json:"company_info_id" gorm:"foreign_key:ID"`
	//has many
	Requests 			[]Request
	Matchings 		[]Matching
	WatchLists 		[]WatchList
	//time stamp
	UpdatedAt			time.Time		`json:"updated_at"`
	CreatedAt			time.Time		`json:"created_at"`
}

type Users []User
