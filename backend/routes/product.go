package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/schantz/web/go-api-produtos/backend/controllers"
)

func ProductRoute(r *gin.Engine) {
	r.GET("/product", controllers.AllProducts)
	r.POST("/product", controllers.CreateProduct)
	r.PUT("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)
}
