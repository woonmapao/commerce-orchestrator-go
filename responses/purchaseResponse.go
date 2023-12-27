package responses

import "github.com/gin-gonic/gin"

func PurchaseErrorResponse(errors []string) gin.H {
	return gin.H{
		"status":  "error",
		"message": "Purchase failed",
		"data": gin.H{
			"errors": errors,
		},
	}
}
