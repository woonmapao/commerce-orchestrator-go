package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woonmapao/commerce-orchestrator-go/models"
	"github.com/woonmapao/commerce-orchestrator-go/responses"
)

func PurchaseOrder(c *gin.Context) {

	var purchaseOrder models.PurchaseOrder
	err := c.ShouldBindJSON(&purchaseOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responses.PurchaseErrorResponse([]string{
				"Invalid purchase format",
			}))
	}

	// validate input data
	err = ValidatePurchaseOrderData(&purchaseOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responses.PurchaseErrorResponse([]string{
				err.Error(),
			}))
	}

	// check user existence
	user, err := GetUserByID(purchaseOrder.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			responses.PurchaseErrorResponse([]string{
				"Failed to retrieve user",
				err.Error(),
			}))
	}

	// check product existence and available stock + calculate subtotal
	var subtotal float64
	for _, product := range purchaseOrder.ProductsID {

		// find if product exist
		p, err := GetProductByID(product.ProductID)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				responses.PurchaseErrorResponse([]string{
					"Failed to retrieve product",
					err.Error(),
				}))
		}

		// compare purchase quantity and available stock
		if product.quantity > p.Stock {
			c.JSON(http.StatusBadRequest,
				responses.PurchaseErrorResponse([]string{
					"Insufficient stock for ProductID: " + string(product.ProductID),
					err.Error(),
				}))
			return
		}

		subtotal += float64(product.quantity) * p.Price

		UpdateStock(product.ProductID, product.Quantity)
	}
	totalAmount := subtotal

	// create order
	orderID, err := CreateOrder(purchaseOrder.UserID, totalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			responses.PurchaseErrorResponse([]string{
				"Failed to create order",
				err.Error(),
			}))
		return
	}

	// create order detail
	for _, product := range purchaseOrder.Products {
		productInfo, _ := getProductByID(product.ProductID)
		subtotal := float64(product.Quantity) * productInfo.Price
		CreateOrderDetail(orderID.gorm.Model.ID, product.ProductID, product.Quantity, subtotal)
	}

	// * dont for get to do transaction
	// validate input data
	// check user, product exist
	// get product price, stock
	// compare purchase quantity and available stock
	// calculate subtotal
	// update stock + create multiple order details
	// calculate TotalAmount + create an Order
	// link order details to that Order
	// return response

}
