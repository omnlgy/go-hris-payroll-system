package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
)

type LeaveController struct {
	service domain.LeaveService
}

func NewLeaveController(service domain.LeaveService) *LeaveController {
	return &LeaveController{
		service: service,
	}
}

func (c *LeaveController) RequestLeave(ctx *gin.Context) {
	var body dto.RequestLeaveRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	employeeID := ctx.GetUint("employee_id")
	leave, err := c.service.RequestLeave(employeeID, body.StartDate, body.EndDate, body.Reason)
	if err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Data:    leave,
		Message: "Leave requested successfully",
	})
}

func (c *LeaveController) ApproveLeave(ctx *gin.Context) {
	leaveID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: "Invalid leave ID",
		})
		return
	}
	var body dto.ApproveLeaveRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, dto.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}
	leave, err := c.service.ApproveLeave(uint(leaveID), body.Status)
	if err != nil {
		ctx.JSON(500, dto.InternalServerErrorResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.SuccessResponse{
		Data:    leave,
		Message: "Leave approved successfully",
	})
}
