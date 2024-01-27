package main

import (
	"github.com/Om/database" // Assuming this is the package where InitDB is located
	"github.com/Om/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	db := database.InitDB()
	gin.SetMode(gin.ReleaseMode)

	routes.ProductRoutes(r, db)
	routes.UserRoutes(r, db)


	r.Run(":8080")
}
