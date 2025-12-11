package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/dto"
	"github.com/harshdevops117/service"
)

type RegisterController struct {
	service *service.RegisterService
}

func NewRegisterController(s *service.RegisterService) *RegisterController {
	return &RegisterController{service: s}
}

func (c *RegisterController) Register(ctx *gin.Context) {
	var body dto.RegisterDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.Register(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}
