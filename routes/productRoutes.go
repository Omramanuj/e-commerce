package routes

import (
	"github.com/Om/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(router *gin.Engine, db *gorm.DB) {
	productHandler := handlers.New(db)

	router.GET("/products", productHandler.GetAllProducts)
	router.POST("/addProduct", productHandler.CreateProduct)
	router.GET("/product/:id", productHandler.GetProduct)
	router.PUT("/updateProduct/:id", productHandler.UpdateProduct)
	router.DELETE("/deleteProduct/:id", productHandler.DeleteProduct)
}