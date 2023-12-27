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
