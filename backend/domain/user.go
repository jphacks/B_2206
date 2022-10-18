package domain

type User struct {
	ID             int    `json:"id" gorm:"primary_key"`
	Name           string `json:"name" gorm:"not null"`
	Email          string `json:"email" gorm:"not null"`
	Password       string `json:"password" gorm:"not null"`
	PersonalInfoId int    `json:"personal_info_id" gorm:"not null"`
	CompanyInfoId  int    `json:"company_info_id" gorm:"not null"`
	UpdatedAt      string `json:"updated_at"`
	CreatedAt      string `json:"created_at"`
}

type Users []User
