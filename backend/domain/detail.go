package domain

type Detail struct {
	ID int `json:"id" gorm:"primary_key"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}