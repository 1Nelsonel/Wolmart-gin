package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	database "github.com/1Nelsonel/Wolmart/Database"
	model "github.com/1Nelsonel/Wolmart/Model"
)

// Update createProduct, updateProduct, getProduct, and other functions to use GORM for database operations
// ========================================================================================================

// ===========GET ALL PRODUCTS==========
func GetProducts(c *gin.Context) {
	db := database.DBConn
	var products []model.Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// =======GET A PRODUCT==================
func GetProduct(c *gin.Context) {
	db := database.DBConn
	id := c.Param("id")
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// ========CREATE A PRODUCT================
func CreateProduct(c *gin.Context) {
	db := database.DBConn
	var newProduct model.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newProduct)

	c.JSON(http.StatusCreated, newProduct)
}

// ==========UPDATE A PRODUCT================
func UpdateProduct(c *gin.Context) {
	db := database.DBConn
	id := c.Param("id")
	var updatedProduct model.Product

	if err := db.First(&updatedProduct, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&updatedProduct)

	c.JSON(http.StatusOK, updatedProduct)
}

// ============DELETE PRODUCT =================
func DeleteProduct(c *gin.Context) {
	db := database.DBConn
	id := c.Param("id")
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// ===============GET PRODUCT BY CATEGORY ==========
func GetProductsByCategory(c *gin.Context) {
	db := database.DBConn
	category := c.Param("category")
	var categoryProducts []model.Product
	db.Where("category = ?", category).Find(&categoryProducts)

	c.JSON(http.StatusOK, categoryProducts)
}
