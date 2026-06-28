package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/controller"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"github.com/omnlgy/go-hris-payroll-system/internal/repository"
	"github.com/omnlgy/go-hris-payroll-system/internal/router"
	"github.com/omnlgy/go-hris-payroll-system/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=admin password=admin123 dbname=hris port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Department{},
		&models.Position{},
		&models.Employee{},
		&models.Attendance{},
		&models.Leave{},
		&models.Salary{},
		&models.BlacklistedToken{},
		&models.DepartmentBudget{},
	); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// Repositories
	deptRepo := repository.NewDepartmentRepository(db)
	posRepo := repository.NewPositionRepository(db)
	empRepo := repository.NewEmployeeRepository(db)
	attRepo := repository.NewAttendanceRepository(db)
	leaveRepo := repository.NewLeaveRepository(db)
	salRepo := repository.NewSalaryRepository(db)

	// Services
	deptSvc := service.NewDepartmentService(deptRepo)
	posSvc := service.NewPositionService(posRepo)
	empSvc := service.NewEmployeeService(empRepo, deptSvc, posSvc)
	attSvc := service.NewAttendanceService(attRepo)
	leaveSvc := service.NewLeaveService(leaveRepo)
	salSvc := service.NewSalaryService(salRepo, attRepo, empRepo)

	// Controllers
	deptCtrl := controller.NewDepartmentController(deptSvc)
	posCtrl := controller.NewPositionController(posSvc)
	empCtrl := controller.NewEmployeeController(empSvc)
	attCtrl := controller.NewAttendanceController(attSvc)
	leaveCtrl := controller.NewLeaveController(leaveSvc)
	salCtrl := controller.NewSalaryController(salSvc)

	r := gin.Default()
	router.DepartmentRoutes(r, deptCtrl)
	router.PositionRoutes(r, posCtrl)
	router.EmployeeRoutes(r, empCtrl)
	router.AttendanceRoutes(r, attCtrl)
	router.LeaveRoutes(r, leaveCtrl)
	router.SalaryRoutes(r, salCtrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
