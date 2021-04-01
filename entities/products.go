package entities

type Product struct {
	ProductID string
	Provider  int     `json:"provider,omitempty"`
	Link      string  `json:"link,omitempty"`
	Brand     string  `json:"brand,omitempty"`
	Name      string  `json:"name"`
	Stock     bool    `json:"stock"`
	Category  string  `json:"category"`
	Price     float64 `json:"price"`
}
