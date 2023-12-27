package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/commerce-orchestrator-go/models"
	"github.com/woonmapao/commerce-orchestrator-go/responses"
	"github.com/woonmapao/commerce-orchestrator-go/services"
	"github.com/woonmapao/commerce-orchestrator-go/validation"
)

func PurchaseOrder(c *gin.Context) {

	var purchaseOrder models.PurchaseOrder
	err := c.ShouldBindJSON(&purchaseOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responses.PurchaseErrorResponse([]string{
				"Invalid purchase format",
				err.Error(),
			}))
	}

	// validate input data
	err = validation.ValidatePurchaseOrderData(&purchaseOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responses.PurchaseErrorResponse([]string{
				err.Error(),
			}))
	}

	// check user existence
	user, err := services.GetUserByID(purchaseOrder.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responses.PurchaseErrorResponse([]string{
				"Failed to retrieve user",
				err.Error(),
			}))
	}

	// check product existence and available stock + calculate subtotal
	var subtotal float64
	for _, product := range purchaseOrder.Products {

		// find if product exist
		p, err := services.GetProductByID(product.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				responses.PurchaseErrorResponse([]string{
					"Failed to retrieve product",
					err.Error(),
				}))
		}

		// compare purchase quantity and available stock
		if product.Quantity > p.StockQuantity {
			c.JSON(http.StatusBadRequest,
				responses.PurchaseErrorResponse([]string{
					"Insufficient stock for ProductID: " + fmt.Sprint(product.ProductID),
					err.Error(),
				}))
			return
		}

		_, err = services.UpdateStock(product.ProductID, product.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				responses.PurchaseErrorResponse([]string{
					"Failed to update stock",
					err.Error(),
				}))
		}

		subtotal += float64(product.Quantity) * p.Price
		log.Printf("Calculating subtotal: %v + %s:%v", subtotal, p.Name, p.Price)
	}
	totalAmount := subtotal
	log.Printf("Total amount : %v", totalAmount)

	// create order
	orderID, err := services.CreateOrder(purchaseOrder.UserID, totalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.PurchaseErrorResponse([]string{
				"Failed to create order",
				err.Error(),
			}))
		log.Printf("Failed to create order. Error: %s", err.Error())
		return
	}
	log.Printf("Order created successfully, OrderID: %d", orderID)

	// create order detail
	productList := make([]string, len(purchaseOrder.Products))
	for _, product := range purchaseOrder.Products {
		productInfo, _ := services.GetProductByID(product.ProductID)
		subtotal := float64(product.Quantity) * productInfo.Price
		services.CreateOrderDetail(int(orderID), product.ProductID, product.Quantity, subtotal)
		productList = append(productList, productInfo.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Purchase order successful",
		"data": gin.H{
			"userID":            user.ID,
			"username":          user.Username,
			"orderID":           orderID,
			"totalAmount":       totalAmount,
			"product_purchased": productList,
		},
	})

}
