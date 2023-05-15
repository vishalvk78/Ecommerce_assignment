package routes

import (
	"github.com/gin-gonic/gin"

	"ecommerce/controllers"
	"ecommerce/middleware"
)

func ProductRoutes(incomingRoutes *gin.Engine) {

	// Routes accessible by all users
	incomingRoutes.GET("/products/search", controllers.SearchProducts)
	incomingRoutes.GET("/product/:id", controllers.GetProductByID)
	incomingRoutes.GET("/products", controllers.GetProducts)

	// Routes accessible only by admin users
	adminRoutes := incomingRoutes.Group("/admin")
	adminRoutes.Use(middleware.AdminMiddleware())
	{
		adminRoutes.POST("/product", controllers.AddProduct)
		adminRoutes.PATCH("/product/:id", controllers.UpdateProductDetails)
		adminRoutes.DELETE("/product/:id", controllers.DeleteProduct)
	}

}
