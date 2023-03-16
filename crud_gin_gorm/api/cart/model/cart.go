package model

type Cart struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
}

type CartItem struct {
	ID        uint    `json:"id"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}
