package main

import (
	"github.com/Om/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.ProductRoutes(r) 
	// r.GET("/products", h.GetAllProducts)
	// r.POST("/createProduct",h.CreateProduct)
	r.Run(":8080")
}
