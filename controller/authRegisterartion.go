package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/dto"
	"github.com/harshdevops117/models"
	"github.com/harshdevops117/service"
)

type AuthController struct {
	registerService *service.RegisterUser
	loginService    *service.LoginService
}

func NewAuthController(
	registerService *service.RegisterUser,
	loginService *service.LoginService,
) *AuthController {
	return &AuthController{
		registerService: registerService,
		loginService:    loginService,
	}
}

func (c *AuthController) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")

	// REGISTER
	auth.POST("/register", func(ctx *gin.Context) {
		var user models.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request body",
			})
			return
		}

		if err := c.registerService.RegisterUser(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "registered successfully",
		})
	})

	// LOGIN
	auth.POST("/login", func(ctx *gin.Context) {
		var req dto.LoginDTO

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request body",
			})
			return
		}

		user, err := c.loginService.Login(req.Email, req.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"user": gin.H{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
		})
	})
}
