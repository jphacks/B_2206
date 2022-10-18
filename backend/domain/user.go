package domain

import (
	"time"
)

type User struct {
	ID				int    			`json:"id" gorm:"primary_key"`
	Name			string			`json:"name" gorm:"not null"`
	Email			string			`json:"email" gorm:"not null"`
	Password		string			`json:"password" gorm:"not null"`
	PersonalInfo 	PersonalInfo
	CompanyInfo  	CompanyInfo
	Requests 		[]Request
	Matchings 		[]Matching
	WatchLists 		[]WatchList
	UpdatedAt		time.Time		`json:"updated_at"`
	CreatedAt		time.Time		`json:"created_at"`
}

type Users []User
