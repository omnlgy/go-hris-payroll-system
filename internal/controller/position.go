package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"github.com/omnlgy/go-hris-payroll-system/internal/repository"
)

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
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Positions retrieved successfully",
		"data":    positions,
	})
}

func (c *PositionController) CreatePosition(ctx *gin.Context) {
	var body dto.CreatePositionRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid request body",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "body",
					Message: err.Error(),
				},
			},
		})
		return
	}

	position := &models.Position{
		Title:      body.Title,
		BaseSalary: body.BaseSalary,
	}

	createdPosition, err := c.service.CreatePosition(position)

	if err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(201, dto.SuccessResponse{
		Message: "Position created successfully",
		Data:    createdPosition,
	})
}

func (c *PositionController) UpdatePosition(ctx *gin.Context) {
	var body dto.CreatePositionRequest
	var PositionID uint64

	PositionID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid request body",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "id",
					Message: err.Error(),
				},
			},
		})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid request body",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "body",
					Message: err.Error(),
				},
			},
		})
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
			ctx.JSON(404, dto.NotFoundResponse{
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Position updated successfully",
		Data:    updatedPosition,
	})
}

func (c *PositionController) DeletePosition(ctx *gin.Context) {
	positionID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid position ID",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "id",
					Message: err.Error(),
				},
			},
		})
		return
	}

	err = c.service.DeletePosition(uint(positionID))
	if err != nil {
		if errors.Is(err, repository.PositionNotFound) {
			ctx.JSON(404, dto.NotFoundResponse{
				Message: err.Error(),
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Position deleted successfully",
	})
}
