package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/controller"
)

func RegisterRoutes(r *gin.RouterGroup, controller *controller.RegisterController) {
	r.POST("/register", controller.Register)
}
