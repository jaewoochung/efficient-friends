package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	// add in the imports once you have setup the routes directory and files
	"server/routes"
)

func main() {
	// Initialize the ports
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// gin is a web framework for Go - helps for API management
	router := gin.New()
	router.Use(gin.Logger())

	// Need to resolve the cors setup for the router...
	router.Use(cors.Default())

	// Setup the endpoints
	/*
	 * Example of a route
	 * router.POST("/order/create", routes.AddOrder)
	 */
	router.POST("/question/create", routes.AddQuestion)

	// run the server and allow it to listen to requests
	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
