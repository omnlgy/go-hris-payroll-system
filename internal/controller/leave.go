package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
)

type RequestLeaveRequest struct {
	EmployeeID uint   `json:"employeeId" binding:"required"`
	StartDate  string `json:"startDate" binding:"required"`
	EndDate    string `json:"endDate" binding:"required"`
	Reason     string `json:"reason" binding:"required"`
}

type ApproveLeaveRequest struct {
	Status string `json:"status" binding:"required"`
}

type LeaveController struct {
	service domain.LeaveService
}

func NewLeaveController(service domain.LeaveService) *LeaveController {
	return &LeaveController{
		service: service,
	}
}

func (c *LeaveController) RequestLeave(ctx *gin.Context) {
	var body RequestLeaveRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	leave, err := c.service.RequestLeave(body.EmployeeID, body.StartDate, body.EndDate, body.Reason)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"data":    leave,
		"message": "Leave requested successfully",
	})
}

func (c *LeaveController) ApproveLeave(ctx *gin.Context) {
	leaveID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid leave ID"})
		return
	}
	var body ApproveLeaveRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	leave, err := c.service.ApproveLeave(uint(leaveID), body.Status)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"data":    leave,
		"message": "Leave approved successfully",
	})
}
