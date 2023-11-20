package presentation

type (
	Product struct {
		Link        string `json:"link"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Rating      string `json:"rating"`
		Merchant    string `json:"merchant"`
	}

	Products struct {
		Products []Product `json:"products"`
	}
)
