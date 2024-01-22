package routes

import (
	"github.com/Om/database"
	"github.com/Om/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.Engine) {
	DB := database.InitDB()
	h := handlers.New(DB)
	User := c.Group("/users")
	{
		User.GET("/Users", h.GetAllUsers)
		User.GET("/User/:id", h.GetUser)
		User.POST("/AddUser", h.AddUser)
		User.PUT("/Users", h.UpdateUser)

		// Product.POST("/CreateUser",h.CreateUser)
	}
}
