package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/controller"
	database "github.com/harshdevops117/db"
	"github.com/harshdevops117/models"
	"github.com/harshdevops117/service"
)

func main() {
	app := gin.Default()

	app.SetTrustedProxies([]string{"127.0.0.1"})

	dbConn, err := database.DataBaseInit()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := dbConn.AutoMigrate(&models.User{}, &models.Notes{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Cats",
		})
	})

	registerService := service.NewRegisterUser(dbConn)
	loginService := service.NewLoginService(dbConn)
	notesService := service.NewNotesService(dbConn)

	authController := controller.NewAuthController(registerService, loginService)
	notesController := controller.NewNotesController(notesService)

	authController.RegisterRoutes(app)
	notesController.RegisterRoutes(app)

	if err := app.Run(":8000"); err != nil {
		log.Fatal("server failed to start:", err)
	}
}
