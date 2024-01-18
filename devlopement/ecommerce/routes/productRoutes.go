package routes

import (
	"github.com/Om/database"
	"github.com/Om/handlers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(c *gin.Engine) {
	DB := database.InitDB()
	h := handlers.New(DB)
	Product := c.Group("/products")
	{
		Product.GET("/list", h.GetAllProducts)
		Product.GET("/product/:id", h.GetProduct)
		Product.POST("/createProduct", h.CreateProduct)
		Product.PUT("/updateProduct", h.UpdateProduct)
	}
}
