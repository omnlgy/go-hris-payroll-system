package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
)

type SalaryController struct {
	service domain.SalaryService
}

func NewSalaryController(service domain.SalaryService) *SalaryController {
	return &SalaryController{
		service: service,
	}
}

func (c *SalaryController) CalculateSalary(ctx *gin.Context) {
	var body dto.SalaryRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	_, err := time.Parse("2006-01", body.Period)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid period format, expected YYYY-MM",
		})
		return
	}

	calculatedSalary, err := c.service.CalculateSalary(body.EmployeeID, body.Period)
	if err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, dto.SuccessResponse{
		Message: "Salary calculated successfully",
		Data:    calculatedSalary,
	})
}

func (c *SalaryController) GetSaleryEmployeeByPeriod(ctx *gin.Context) {
	period := ctx.Param("period")
	employeeID, err := strconv.ParseUint(ctx.Query("employee_id"), 10, 64)

	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid employee ID",
		})
		return
	}

	_, err = time.Parse("2006-01", period)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid period format, expected YYYY-MM",
		})
		return
	}
	salaries, err := c.service.GetSaleryEmployeeByPeriod(uint(employeeID), period)
	if err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, dto.SuccessResponse{
		Message: "Salaries retrieved successfully",
		Data:    salaries,
	})
}
