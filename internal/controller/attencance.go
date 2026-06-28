package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
)

type RecordAttendanceRequest struct {
	EmployeeID uint `json:"employee_id" binding:"required"`
}

type AttendanceController struct {
	service domain.AttendanceService
}

func NewAttendanceController(service domain.AttendanceService) *AttendanceController {
	return &AttendanceController{
		service: service,
	}
}

func (c *AttendanceController) RecordAttendance(ctx *gin.Context) {
	var body RecordAttendanceRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Attend(body.EmployeeID); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Attendance recorded successfully"})
}
