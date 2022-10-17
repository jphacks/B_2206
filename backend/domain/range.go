package domain

type Range struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	MaxValueID int `json:"max_value_id" gorm:"foreign_key:ID"`
	MinValueID int `json:"min_value_id" gorm:"foreign_key:ID"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}