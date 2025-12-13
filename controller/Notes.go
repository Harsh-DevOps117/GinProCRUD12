package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harshdevops117/dto"
	"github.com/harshdevops117/service"
)

type NotesController struct {
	service *service.NotesService
}

func NewNotesController(service *service.NotesService) *NotesController {
	return &NotesController{service: service}
}

func (c *NotesController) RegisterRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	{
		notes.POST("", c.Create)
		notes.GET("", c.GetAll)
		notes.GET("/:id", c.GetOne)
		notes.PUT("/:id", c.Update)
		notes.DELETE("/:id", c.Delete)
	}
}

func getUserID(ctx *gin.Context) uint {
	return 1
}

func (c *NotesController) Create(ctx *gin.Context) {
	var req dto.NotesDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := c.service.CreateNote(getUserID(ctx), req.Title, req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "note created"})
}

func (c *NotesController) GetAll(ctx *gin.Context) {
	notes, err := c.service.GetNotes(getUserID(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

func (c *NotesController) GetOne(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	note, err := c.service.GetNoteByID(getUserID(ctx), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}

func (c *NotesController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var req dto.NotesDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := c.service.UpdateNote(
		getUserID(ctx),
		uint(id),
		req.Title,
		req.Content,
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "note updated"})
}

func (c *NotesController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.DeleteNote(getUserID(ctx), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "note deleted"})
}
