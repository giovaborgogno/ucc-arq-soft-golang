package items

type Item struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	HasTaxes    bool    `json:"has_taxes"`
}

const ItemStore = "Nike tienda oficial"

func GetItem() *Item {
	return &Item{
		Title:       "Zapatillas",
		Description: "Nuevas zapatillas",
		Price:       30000,
		HasTaxes:    false,
	}
}
