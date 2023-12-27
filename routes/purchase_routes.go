package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/woonmapao/commerce-orchestrator-go/handlers"
)

func SetUpPurchaseRoutes(router *gin.Engine) {
	purchaseGroup := router.Group("/purchase")
	{
		purchaseGroup.POST("/", handlers.PurchaseOrder)
	}
}
