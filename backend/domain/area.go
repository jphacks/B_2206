package domain

type Area struct {
	ID int `json:"id" gorm:"primary_key"`
	PostCode int `json:"post_code" gorm:"not null"`
	Prefecture string `json:"prefecture" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	AddressNumber int `json:"address_number"`
	BuildingName string `json:"building_name"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}