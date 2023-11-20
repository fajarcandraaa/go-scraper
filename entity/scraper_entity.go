package entity

type Product struct {
	Link        string `json:"link" gorm:"type:text"`
	Name        string `json:"name" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
	Price       string `json:"price" gorm:"size:255"`
	Rating      string `json:"rating" gorm:"size:255"`
	Merchant    string `json:"merchant" gorm:"size:255"`
}
