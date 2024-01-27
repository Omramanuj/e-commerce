package handlers

import (
	"github.com/Om/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




type handler struct {
    DB *gorm.DB
}


func New(db *gorm.DB) handler {
    return handler{DB: db} // Set DB field with the provided DB instance
}

func (h handler) GetAllProducts(c *gin.Context) {
	var products []models.Product
	h.DB.Find(&products)
	c.JSON(200, products)
}

func (h handler) CreateProduct(c *gin.Context){
	var p models.Product
	if err:=c.ShouldBindJSON(&p); err!= nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	if err := h.DB.Create(&p).Error; err!= nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200, p)
}

func (h handler) GetProduct(c *gin.Context){
	var user models.Product
	id := c.Param("id")
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func (h handler) UpdateProduct(c *gin.Context){
	var user models.Product
	if err:= c.ShouldBindJSON(&user);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
	}
	if err:= h.DB.Save(&user).Error; err!=nil{
		c.JSON(500,gin.H{"error":"Failed to update user"})
		return
	}
	c.JSON(200,user)
}


func (h handler) DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product

    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Product not found"})
        return
    }
    if err := h.DB.Delete(&product).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete product"})
        return
    }

    c.JSON(200, gin.H{"message": "Product deleted successfully"})
}