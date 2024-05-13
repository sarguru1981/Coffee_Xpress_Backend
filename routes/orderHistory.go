package routes

import (
	controller "golang-Coffee_Xpress_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderHistoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
