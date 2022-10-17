package domain

type DetailsTag struct {
	ID int `json:"id" gorm:"primary_key"`
	DetailID int `json:"detail_id" gorm:"foreign_key:ID"`
	TagID int `json:"tag_id" gorm:"foreign_key:ID"`
	ClassificationID string `json:"classification_id" gorm:"foreign_key:ID"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}