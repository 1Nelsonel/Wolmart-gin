package route

import (
	"github.com/gin-gonic/gin"

	controller "github.com/1Nelsonel/Wolmart/Controller"
)

func SetupRoutes(r *gin.Engine) {
	// Define routes
	r.GET("/products", controller.GetProduct)
	r.GET("/products/:id", controller.GetProduct)
	r.POST("/products", controller.CreateProduct)
	r.PUT("/products/:id", controller.UpdateProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)
	r.GET("/products/category/:category", controller.GetProductsByCategory)

}
