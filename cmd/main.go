package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/controller"
	"github.com/harshdevops117/db"
	"github.com/harshdevops117/models"
	"github.com/harshdevops117/service"
)

func main() {
	app := gin.Default()

	db,err:=db.DataBaseInit()
	if err!=nil{
		panic(err)
	}
	db.AutoMigrate(&models.User{},&models.Notes{})

	app.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"Cats",
		})
	})

	serviceRegistration:=service.NewRegisterUser(db)
	serviceRegistration1:=service.NewLoginService(db)
	ControllerRegistration:=controller.NewAuthController(serviceRegistration,serviceRegistration1)

	ControllerRegistration.RegisterRoutes(app)

	app.Run(":8000")
}
