package domain

type Value struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Value string `json:"value"`
	RangeID int `json:"range_id" gorm:"foreign_key:ID"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

