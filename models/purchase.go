package models

type PurchaseOrder struct {
	UserID      uint      `json:"userId" binding:"required"`
	Products    []Product `json:"products" binding:"required"`
	TotalAmount float64   `json:"totalAmount" binding:"required,gte=1"`
}
