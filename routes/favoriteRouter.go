package routes

import (
	controller "golang-Coffee_Xpress_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/favorites", controller.GetFavorites())
	incomingRoutes.GET("/favorites/:favorite_id", controller.GetFavorite())
	incomingRoutes.POST("/favorites", controller.CreateFavorite())
	incomingRoutes.PATCH("/favorites/:favorite_id", controller.UpdateFavorite())
}
