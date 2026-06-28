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

type EmployeeController struct {
	service domain.EmployeeService
}

func NewEmployeeController(service domain.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var body dto.CreateEmployeeRequest

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
			ctx.JSON(400, dto.BadRequestResponse{
				Message: "Department or Position not found",
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(201, dto.SuccessResponse{
		Message: "Employee created successfully",
		Data:    employee,
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
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Employees retrieved successfully",
		Data:    employees,
	})
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	employeeID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid employee ID",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "id",
					Message: "Invalid employee ID",
				},
			},
		})
		return
	}

	err = c.service.DeleteEmployee(uint(employeeID))
	if err != nil {
		if errors.Is(err, repository.EmployeeNotFound) {
			ctx.JSON(404, dto.NotFoundResponse{
				Message: "Employee not found",
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Employee deleted successfully",
	})
}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	employeeID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid employee ID",
			Errors: []struct {
				Field   string `json:"field"`
				Message string `json:"message"`
			}{
				{
					Field:   "id",
					Message: "Invalid employee ID",
				},
			},
		})
		return
	}

	var body dto.UpdateEmployeeRequest

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

	employee := &models.Employee{
		ID:           uint(employeeID),
		NIK:          body.NIK,
		FullName:     body.FullName,
		Email:        body.Email,
		DepartmentID: body.DepartmentID,
		PositionID:   body.PositionID,
		Role:         body.Role,
		Status:       body.Status,
	}
	if body.Password != "" {
		employee.Password = body.Password
	}

	if _, err := c.service.Update(employee); err != nil {
		if errors.Is(err, repository.EmployeeNotFound) {
			ctx.JSON(404, dto.NotFoundResponse{
				Message: "Employee not found",
			})
			return
		}
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Employee updated successfully",
		Data:    employee,
	})
}
