package models

import "gorm.io/gorm"

type PurchaseProduct struct {
	ProductID int `json:"productID" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,gte=1"`
}

type Product struct {
	gorm.Model
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	StockQuantity int     `json:"stockQuantity"`
	ReorderLevel  int     `json:"reorderLevel"`
}
