package models

type Product struct {
	ProductID int `json:"productID" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,gte=1"`
}
