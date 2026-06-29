package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/controller"
	"github.com/omnlgy/go-hris-payroll-system/internal/middleware"
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
	apiDepartments.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("HRD"), controller.DeletePosition)
}

func EmployeeRoutes(router *gin.Engine, controller *controller.EmployeeController) {
	apiEmployees := router.Group("/api/employees")

	apiEmployees.POST("/", middleware.AuthMiddleware(), middleware.RequireRole("HRD"), controller.CreateEmployee)
	apiEmployees.GET("/", controller.GetEmployees)
	apiEmployees.DELETE("/:id", middleware.AuthMiddleware(), middleware.RequireRole("HRD"), controller.DeleteEmployee)
	apiEmployees.PUT("/:id", middleware.AuthMiddleware(), middleware.RequireRole("HRD"), controller.UpdateEmployee)
}

func AttendanceRoutes(router *gin.Engine, controller *controller.AttendanceController) {
	api := router.Group("/api")

	api.POST("/attendance", middleware.AuthMiddleware(), controller.RecordAttendance)
}

func LeaveRoutes(router *gin.Engine, controller *controller.LeaveController) {
	api := router.Group("/api")

	api.POST("/leaves", middleware.AuthMiddleware(), controller.RequestLeave)
	api.PATCH("/leaves/:id/approve", middleware.AuthMiddleware(), middleware.RequireRole("HRD"), controller.ApproveLeave)
}

func SalaryRoutes(router *gin.Engine, controller *controller.SalaryController) {
	api := router.Group("/api/salaries")

	api.GET("/period/:period", controller.GetSaleryEmployeeByPeriod)
	api.POST("/calculate", controller.CalculateSalary)
}

func AuthRoutes(router *gin.Engine, controller *controller.AuthController) {
	api := router.Group("/api/auth")

	api.POST("/login", controller.Login)
	api.POST("/register", controller.Register)
	api.POST("/logout", controller.Logout)
}
