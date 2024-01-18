package handlers

import (
	"github.com/Om/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllUsers(c *gin.Context) {
	var users []models.User
	h.DB.Find(&users)
	c.JSON(200, users)
}

func (h handler) AddUser(c *gin.Context) {
	var user []models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Create a new user record in the database
	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, gin.H{"message": "User created successfully", "user": user})

}

func (h handler) GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func (h handler) UpdateUser(c *gin.Context) {
	var user []models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err:= h.DB.Save(&user).Error; err!= nil{
		c.JSON(500,gin.H{"error":"Failed to update user"})
		return
	}
	c.JSON(201, gin.H{"message": "User updated successfully", "user": user})
}