package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
)

type AttendanceController struct {
	service domain.AttendanceService
}

func NewAttendanceController(service domain.AttendanceService) *AttendanceController {
	return &AttendanceController{
		service: service,
	}
}

func (c *AttendanceController) RecordAttendance(ctx *gin.Context) {
	employeeID := ctx.GetUint("employee_id")

	if err := c.service.Attend(employeeID); err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Message: "Attendance recorded successfully",
	})
}
