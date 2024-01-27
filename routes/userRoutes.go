package routes

import (
	"github.com/Om/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	userHandler := handlers.New(db)

	router.GET("/users", userHandler.GetAllUsers)
	router.POST("/addUser", userHandler.AddUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

}
