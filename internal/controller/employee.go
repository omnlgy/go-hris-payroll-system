package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"github.com/omnlgy/go-hris-payroll-system/internal/repository"
)

type CreateEmployeeRequest struct {
	NIK          string `json:"nik" binding:"required"`
	FullName     string `json:"fullName" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID uint   `json:"departmentId" binding:"required"`
	PositionID   uint   `json:"positionId" binding:"required"`
	Role         string `json:"role" binding:"required,oneof=ADMIN MANAGER EMPLOYEE"`
	Password     string `json:"password" binding:"required,min=6"`
	Status       string `json:"status" binding:"required,oneof=ACTIVE SUSPENDED TERMINATED"`
}

type EmployeeController struct {
	service domain.EmployeeService
}

func NewEmployeeController(service domain.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var body CreateEmployeeRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	employee := &models.Employee{
		NIK:          body.NIK,
		FullName:     body.FullName,
		Email:        body.Email,
		DepartmentID: body.DepartmentID,
		PositionID:   body.PositionID,
		Role:         body.Role,
		Password:     body.Password,
		Status:       body.Status,
	}

	if _, err := c.service.Add(employee); err != nil {
		if errors.Is(err, repository.DepartmentNotFound) || errors.Is(err, repository.PositionNotFound) {
			ctx.JSON(400, gin.H{"error": "Department or Position not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Employee created successfully",
		"data":    employee,
	})
}

func (c *EmployeeController) GetEmployees(ctx *gin.Context) {
	departmentID, _ := strconv.ParseUint(ctx.Query("departmentId"), 10, 64)
	positionID, _ := strconv.ParseUint(ctx.Query("positionId"), 10, 64)

	filter := models.FilterEmployee{
		Name:         ctx.Query("search"),
		Status:       ctx.Query("status"),
		DepartmentID: uint(departmentID),
		PositionID:   uint(positionID),
	}

	employees, err := c.service.GetEmployees(filter)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Employees retrieved successfully",
		"data":    employees,
		"query":   filter,
	})
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	employeeID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid employee ID"})
		return
	}

	err = c.service.DeleteEmployee(uint(employeeID))
	if err != nil {
		if errors.Is(err, repository.EmployeeNotFound) {
			ctx.JSON(404, gin.H{"error": "Employee not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Employee deleted successfully",
	})
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	employeeID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid employee ID"})
		return
	}

	var body CreateEmployeeRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	employee := &models.Employee{
		ID:           uint(employeeID),
		NIK:          body.NIK,
		FullName:     body.FullName,
		Email:        body.Email,
		DepartmentID: body.DepartmentID,
		PositionID:   body.PositionID,
		Role:         body.Role,
		Password:     body.Password,
		Status:       body.Status,
	}

	if _, err := c.service.Update(employee); err != nil {
		if errors.Is(err, repository.EmployeeNotFound) {
			ctx.JSON(404, gin.H{"error": "Employee not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Employee updated successfully",
		"data":    employee,
	})
}