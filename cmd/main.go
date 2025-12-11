package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/controller"
	"github.com/harshdevops117/db"
	"github.com/harshdevops117/middleware"
	"github.com/harshdevops117/routes"
	"github.com/harshdevops117/service"
)

func main() {
	database, err := db.DataBaseInit()
	if err != nil {
		log.Fatal(err)
	}

	db.DBAutoMigrate(database)

	// -------------------------
	// Initialize Services
	// -------------------------
	registerService := service.NewRegisterService(database)

	// -------------------------
	// Initialize Controllers
	// -------------------------
	registerController := controller.NewRegisterController(registerService)

	// -------------------------
	// Initialize Gin App
	// -------------------------
	app := gin.Default()

	app.Use(middleware.RequestLogger())

	// -------------------------
	// Base route
	// -------------------------
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Good Day By GIN",
		})
	})

	// -------------------------
	// API Routes
	// -------------------------
	api := app.Group("/api")
	auth := api.Group("/auth")

	// attach routes from routes folder
	routes.RegisterRoutes(auth, registerController)

	// -------------------------
	// Start server
	// -------------------------
	app.Run(":8000")
}
