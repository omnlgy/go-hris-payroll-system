package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/controller"
)

func DepartmentRoutes(router *gin.Engine, controller *controller.DepartmentController) {
	apiDepartments := router.Group("/api/departments")

	apiDepartments.GET("/", controller.GetDepartments)
	apiDepartments.POST("/", controller.CreateDepartment)
	apiDepartments.PUT("/:id", controller.UpdateDepartment)
	apiDepartments.DELETE("/:id", controller.DeleteDepartment)
}

func PositionRoutes(router *gin.Engine, controller *controller.PositionController) {
	apiDepartments := router.Group("/api/positions")

	apiDepartments.GET("/", controller.GetPositions)
	apiDepartments.POST("/", controller.CreatePosition)
	apiDepartments.PUT("/:id", controller.UpdatePosition)
	apiDepartments.DELETE("/:id", controller.DeletePosition)
}

func EmployeeRoutes(router *gin.Engine, controller *controller.EmployeeController) {
	apiEmployees := router.Group("/api/employees")

	apiEmployees.POST("/", controller.CreateEmployee)
	apiEmployees.GET("/", controller.GetEmployees)
	apiEmployees.DELETE("/:id", controller.DeleteEmployee)
	apiEmployees.PUT("/:id", controller.UpdateEmployee)
}

func AttendanceRoutes(router *gin.Engine, controller *controller.AttendanceController) {
	api := router.Group("/api")

	api.POST("/attendance", controller.RecordAttendance)
}

func LeaveRoutes(router *gin.Engine, controller *controller.LeaveController) {
	api := router.Group("/api")

	api.POST("/leaves", controller.RequestLeave)
	api.PATCH("/leaves/:id/approve", controller.ApproveLeave)
}

func SalaryRoutes(router *gin.Engine, controller *controller.SalaryController) {
	api := router.Group("/api/salaries")

	api.GET("/period/:period", controller.GetSaleryEmployeeByPeriod)
	api.POST("/calculate", controller.CalculateSalary)
}
