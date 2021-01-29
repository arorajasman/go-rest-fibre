package model

// Product struct
type Product struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

// Products struct
type Products struct {
	Products []Product `json:"products"`
}
