package routes

import (
	controller "golang-Coffee_Xpress_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/products", controller.GetProducts())
	incomingRoutes.GET("/products/:product_id", controller.GetProduct())
	incomingRoutes.POST("/products", controller.CreateProduct())
	incomingRoutes.PATCH("/products/:product_id", controller.UpdateProduct())
}
