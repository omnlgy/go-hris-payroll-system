package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"github.com/omnlgy/go-hris-payroll-system/internal/repository"
)

type CreatePositionRequest struct {
	Title      string  `json:"title" binding:"required"`
	BaseSalary float64 `json:"baseSalary" binding:"required"`
}

type PositionController struct {
	service domain.PositionService
}

func NewPositionController(service domain.PositionService) *PositionController {
	return &PositionController{
		service: service,
	}
}

func (c *PositionController) GetPositions(ctx *gin.Context) {
	positions, err := c.service.GetPositions()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Positions retrieved successfully",
		"data":    positions,
	})
}

func (c *PositionController) CreatePosition(ctx *gin.Context) {
	var body CreatePositionRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	position := &models.Position{
		Title:      body.Title,
		BaseSalary: body.BaseSalary,
	}

	createdPosition, err := c.service.CreatePosition(position)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Position created successfully",
		"data":    createdPosition,
	})
}

func (c *PositionController) UpdatePosition(ctx *gin.Context) {
	var body CreatePositionRequest
	var PositionID uint64

	PositionID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid position ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	position := &models.Position{
		ID:         uint(PositionID),
		Title:      body.Title,
		BaseSalary: body.BaseSalary,
	}

	updatedPosition, err := c.service.UpdatePosition(position)

	if err != nil {
		if errors.Is(err, repository.PositionNotFound) {
			ctx.JSON(404, gin.H{"error": "Position not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Position updated successfully",
		"data":    updatedPosition,
	})
}

func (c *PositionController) DeletePosition(ctx *gin.Context) {
	positionID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid position ID"})
		return
	}

	err = c.service.DeletePosition(uint(positionID))
	if err != nil {
		if errors.Is(err, repository.PositionNotFound) {
			ctx.JSON(404, gin.H{"error": "Position not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Position deleted successfully",
	})
}
