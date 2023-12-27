package main

import (
	"os"

	"github.com/gin-gonic/gin"
	initializer "github.com/woonmapao/commerce-orchestrator-go/Initializer"
	"github.com/woonmapao/commerce-orchestrator-go/routes"
)

func init() {
	initializer.LoadEnvVariables()
}

func main() {

	r := gin.Default()

	routes.SetUpPurchaseRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run()
}
