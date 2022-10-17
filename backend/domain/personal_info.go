package domain

type PersonalInfo struct {
	ID int `json:"id" gorm:"primary_key"`
	FamilyName string `json:"family_name" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	Birthday string `json:"birthday" gorm:"not null"`
	PhoneNumber int `json:"phone_number" gorm:"not null"`
	UserID int `json:"user_id" gorm:"foreign_key:ID"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}