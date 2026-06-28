package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"github.com/omnlgy/go-hris-payroll-system/internal/repository"
)

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

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
	var body CreateDepartmentRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	department := &models.Department{
		Name: body.Name,
		Code: body.Code,
	}

	createdDepartment, err := c.service.CreateDepartment(department)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Department created successfully",
		"data":    createdDepartment,
	})
}

func (c *DepartmentController) UpdateDepartment(ctx *gin.Context) {
	var body CreateDepartmentRequest
	var departmentID uint64

	departmentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid department ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	department := &models.Department{
		ID:   uint(departmentID),
		Name: body.Name,
		Code: body.Code,
	}

	updatedDepartment, err := c.service.UpdateDepartment(department)

	if err != nil {
		if errors.Is(err, repository.DepartmentNotFound) {
			ctx.JSON(404, gin.H{"error": "Department not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Department updated successfully",
		"data":    updatedDepartment,
	})
}

func (c *DepartmentController) DeleteDepartment(ctx *gin.Context) {
	departmentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid department ID"})
		return
	}

	err = c.service.DeleteDepartment(uint(departmentID))
	if err != nil {
		if errors.Is(err, repository.DepartmentNotFound) {
			ctx.JSON(404, gin.H{"error": "Department not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Department deleted successfully",
	})
}
