package domain

type Classification struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
