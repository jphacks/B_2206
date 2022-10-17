package domain

type WatchList struct {
	ID int `json:"id" gorm:"primary_key"`
	UserID int `json:"user_id" gorm:"foreign_key:ID"`
	DetailID int `json:"detail_id" gorm:"foreign_key:ID"`
	UpdatedAt string `json:"updated_at"`
	CreatedAts tring `json:"created_at"`
}