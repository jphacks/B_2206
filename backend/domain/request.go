package domain

type Request struct {
	ID int `json:"id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null"`
	IsPurchase bool `json:"is_purchase" gorm:"not null"`
	UserID int `json:"user_id" gorm:"foreign_key:ID"`
	DetailID int `json:"detail_id" gorm:"foreign_key:ID"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}