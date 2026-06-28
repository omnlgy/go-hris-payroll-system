package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
	"github.com/omnlgy/go-hris-payroll-system/internal/repository"
)

type DepartmentController struct {
	service domain.DepartmentService
}

func NewDepartmentController(service domain.DepartmentService) *DepartmentController {
	return &DepartmentController{
		service: service,
	}
}

func (c *DepartmentController) GetDepartments(ctx *gin.Context) {
	departments, err := c.service.GetDepartments()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Departments retrieved successfully",
		"data":    departments,
	})
}

func (c *DepartmentController) CreateDepartment(ctx *gin.Context) {
	var body dto.CreateDepartmentRequest

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

	input := dto.DepartmentCreate{
		Name: body.Name,
		Code: body.Code,
	}

	createdDepartment, err := c.service.CreateDepartment(input)

	if err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(201, dto.SuccessResponse{
		Message: "Department created successfully",
		Data:    createdDepartment,
	})
}

func (c *DepartmentController) UpdateDepartment(ctx *gin.Context) {
	var body dto.CreateDepartmentRequest
	var departmentID uint64

	departmentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid department ID",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "id",
					Message: "Invalid department ID",
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

	input := dto.DepartmentUpdate{
		ID:   uint(departmentID),
		Name: body.Name,
		Code: body.Code,
	}

	updatedDepartment, err := c.service.UpdateDepartment(input)

	if err != nil {
		if errors.Is(err, repository.DepartmentNotFound) {
			ctx.JSON(404, dto.NotFoundResponse{
				Message: "Department not found",
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Department updated successfully",
		Data:    updatedDepartment,
	})
}

func (c *DepartmentController) DeleteDepartment(ctx *gin.Context) {
	departmentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid department ID",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "id",
					Message: "Invalid department ID",
				},
			},
		})
		return
	}

	err = c.service.DeleteDepartment(uint(departmentID))
	if err != nil {
		if errors.Is(err, repository.DepartmentNotFound) {
			ctx.JSON(404, dto.NotFoundResponse{
				Message: "Department not found",
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Department deleted successfully",
	})
}
