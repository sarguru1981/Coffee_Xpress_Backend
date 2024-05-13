package routes

import (
	controller "golang-Coffee_Xpress_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/payments", controller.GetPayments())
	incomingRoutes.GET("/payments/:payment_id", controller.GetPayment())
	incomingRoutes.POST("/payments", controller.CreatePayment())
	incomingRoutes.PATCH("/payments/:payment_id", controller.UpdatePayment())
}
