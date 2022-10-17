package domain

type Matching struct {
	ID int `json:"id" gorm:"primary_key"`
	BuyerID int `json:"user_id" gorm:"foreign_key:ID"`
	SellerID int `json:"user_id" gorm:"foreign_key:ID"`
	Status string `json:"status"`
	SellerMessage string `json:"seller_message"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}