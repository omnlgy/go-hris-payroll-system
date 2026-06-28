package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
)

type SalaryRequest struct {
	EmployeeID uint   `json:"employee_id" binding:"required"`
	Period     string `json:"period" binding:"required"`
}

type SalaryController struct {
	service domain.SalaryService
}

func NewSalaryController(service domain.SalaryService) *SalaryController {
	return &SalaryController{
		service: service,
	}
}

func (c *SalaryController) CalculateSalary(ctx *gin.Context) {
	var body SalaryRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := time.Parse("2006-01", body.Period)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid period format, expected YYYY-MM"})
		return
	}

	calculatedSalary, err := c.service.CalculateSalary(body.EmployeeID, body.Period)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Salary calculated successfully",
		"data":    calculatedSalary,
	})
}

func (c *SalaryController) GetSaleryEmployeeByPeriod(ctx *gin.Context) {
	period := ctx.Param("period")
	employeeID, err := strconv.ParseUint(ctx.Query("employee_id"), 10, 64)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid employee ID"})
		return
	}

	_, err = time.Parse("2006-01", period)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid period format, expected YYYY-MM"})
		return
	}
	salaries, err := c.service.GetSaleryEmployeeByPeriod(uint(employeeID), period)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Salaries retrieved successfully",
		"data":    salaries,
	})
}
