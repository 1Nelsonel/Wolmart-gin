package route

import (
	"github.com/gin-gonic/gin"

	controller "github.com/1Nelsonel/Wolmart/Controller"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controller.HomePage)
	r.GET("shop/", controller.Shop)
	r.GET("product/", controller.GetProduct)  ///products/:id
	r.POST("products", controller.CreateProduct)
	r.PUT("products/:id", controller.UpdateProduct)
	r.DELETE("products/:id", controller.DeleteProduct)
	r.GET("products/category/:category", controller.GetProductsByCategory)

}
