package main

import (
	"clubApi/config"
	"clubApi/controller"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the database connection
	config.DatabaseInit()

	// Initialize Echo
	e := echo.New()

	// Define a simple root route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// Define routes for users
	userGroup := e.Group("/api/v1/users")
	userGroup.POST("/add", controller.SaveUser)
	userGroup.GET("/", controller.GetAllUser)
	userGroup.GET("/:id", controller.GetUserById)

	// Define routes for clubs
	clubGroup := e.Group("/api/v1/club")
	clubGroup.POST("/add", controller.SaveClub)
	clubGroup.GET("/all", controller.GetAllClubs)
	clubGroup.GET("/:id", controller.GetClubById)
	// Start the server on port 8080 and log any fatal errors
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
