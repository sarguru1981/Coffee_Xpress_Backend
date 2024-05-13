package routes

import (
	controller "golang-Coffee_Xpress_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func CartItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/cartItems", controller.GetCartItems())
	incomingRoutes.GET("/cartItems/:cartItem_id", controller.GetCartItem())
	incomingRoutes.GET("/cartItems-order/:order_id", controller.GetCartItemsByOrder())
	incomingRoutes.POST("/cartItems", controller.CreateCartItem())
	incomingRoutes.PATCH("/cartItems/:cartItem_id", controller.UpdateCartItem())
}
