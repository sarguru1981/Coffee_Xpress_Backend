package main

import (
	"golang-Coffee_Xpress_Backend/database"
	middleware "golang-Coffee_Xpress_Backend/middleware"
	routes "golang-Coffee_Xpress_Backend/routes"

	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "product")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.ProductRoutes(router)
	routes.MenuRoutes(router)
	routes.FavoriteRoutes(router)
	routes.OrderHistoryRoutes(router)
	routes.CartItemRoutes(router)
	routes.PaymentRoutes(router)

	router.Run(":" + port)
}
