package models

type Order struct {
	OrderID     string  `json:"order_id"`
	UserID      int     `json:"user_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	OrderDate   string  `json:"order_date"`
}
