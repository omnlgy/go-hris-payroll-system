package main

import (
	"fmt"
	"log"
	"time"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=admin123 dbname=hris port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Truncate all tables in dependency order for a clean seed
	fmt.Println("Truncating existing data...")
	if err := db.Exec(`
		TRUNCATE TABLE salaries, leaves, attendances, employees, positions, department_budgets, departments, blacklisted_tokens RESTART IDENTITY CASCADE
	`).Error; err != nil {
		log.Fatalf("failed to truncate: %v", err)
	}

	// ---- Departments ----
	depts := []models.Department{
		{Name: "Engineering", Code: "ENG"},
		{Name: "Marketing", Code: "MKT"},
		{Name: "Human Resources", Code: "HR"},
	}
	for i := range depts {
		if err := db.Create(&depts[i]).Error; err != nil {
			log.Fatalf("failed to create department %s: %v", depts[i].Code, err)
		}
		fmt.Printf("  department: %s (id=%d)\n", depts[i].Name, depts[i].ID)
	}

	// ---- Positions ----
	type posDef struct {
		Title      string
		BaseSalary float64
	}
	posData := []posDef{
		{Title: "Engineering Manager", BaseSalary: 20_000_000},
		{Title: "Senior Engineer", BaseSalary: 15_000_000},
		{Title: "Junior Engineer", BaseSalary: 8_000_000},
		{Title: "Marketing Manager", BaseSalary: 15_000_000},
		{Title: "Marketing Specialist", BaseSalary: 7_000_000},
		{Title: "HR Manager", BaseSalary: 12_000_000},
		{Title: "HR Staff", BaseSalary: 6_000_000},
	}
	var positions []models.Position
	for _, p := range posData {
		pos := models.Position{Title: p.Title, BaseSalary: p.BaseSalary}
		if err := db.Create(&pos).Error; err != nil {
			log.Fatalf("failed to create position %s: %v", p.Title, err)
		}
		positions = append(positions, pos)
		fmt.Printf("  position: %s (Rp%.0f, id=%d)\n", pos.Title, pos.BaseSalary, pos.ID)
	}

	// ID-based maps
	posMap := make(map[string]uint)
	for _, p := range positions {
		posMap[p.Title] = p.ID
	}
	deptMap := make(map[string]uint)
	for _, d := range depts {
		deptMap[d.Code] = d.ID
	}

	// ---- Employees ----
	hash := func(pw string) string {
		h, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("hash failed: %v", err)
		}
		return string(h)
	}

	type empDef struct {
		NIK      string
		FullName string
		Email    string
		DeptCode string
		PosTitle string
		Role     string
		Password string
	}

	emps := []empDef{
		{NIK: "001", FullName: "Budi Santoso", Email: "budi@company.com", DeptCode: "ENG", PosTitle: "Engineering Manager", Role: "HRD", Password: "password123"},
		{NIK: "002", FullName: "Siti Rahayu", Email: "siti@company.com", DeptCode: "ENG", PosTitle: "Senior Engineer", Role: "EMPLOYEE", Password: "password123"},
		{NIK: "003", FullName: "Agus Wijaya", Email: "agus@company.com", DeptCode: "ENG", PosTitle: "Junior Engineer", Role: "EMPLOYEE", Password: "password123"},
		{NIK: "004", FullName: "Dewi Lestari", Email: "dewi@company.com", DeptCode: "MKT", PosTitle: "Marketing Manager", Role: "MANAGER", Password: "password123"},
		{NIK: "005", FullName: "Rudi Hartono", Email: "rudi@company.com", DeptCode: "HR", PosTitle: "HR Staff", Role: "HRD", Password: "password123"},
	}

	for _, e := range emps {
		emp := models.Employee{
			NIK:          e.NIK,
			FullName:     e.FullName,
			Email:        e.Email,
			DepartmentID: deptMap[e.DeptCode],
			PositionID:   posMap[e.PosTitle],
			Role:         e.Role,
			Password:     hash(e.Password),
			Status:       "ACTIVE",
		}
		if err := db.Create(&emp).Error; err != nil {
			log.Fatalf("failed to create employee %s: %v", e.NIK, err)
		}
		fmt.Printf("  employee: %s (%s, %s)\n", emp.FullName, emp.Role, emp.Status)
	}

	// ---- Attendance (June 2026, first 3 employees) ----
	fmt.Println("\nSeeding attendance...")
	for empID := uint(1); empID <= 3; empID++ {
		for day := 1; day <= 22; day++ {
			att := models.Attendance{
				EmployeeID: empID,
				Date:       time.Date(2026, time.June, day, 0, 0, 0, 0, time.UTC).Format("2006-01-02"),
				CheckIn:    "09:00",
				CheckOut:   "17:00",
				Status:     "PRESENT",
			}
			if err := db.Create(&att).Error; err != nil {
				log.Fatalf("failed to create attendance emp=%d day=%d: %v", empID, day, err)
			}
		}
	}
	fmt.Println("  attendance: 22 workdays for employees 1-3 (June 2026)")

	// ---- Leaves ----
	fmt.Println("\nSeeding leaves...")
	leaves := []models.Leave{
		{EmployeeID: 2, StartDate: "2026-07-01", EndDate: "2026-07-03", Reason: "Family vacation", Status: "PENDING"},
		{EmployeeID: 3, StartDate: "2026-07-05", EndDate: "2026-07-06", Reason: "Sick leave", Status: "PENDING"},
	}
	for _, l := range leaves {
		if err := db.Create(&l).Error; err != nil {
			log.Fatalf("failed to create leave emp=%d: %v", l.EmployeeID, err)
		}
		fmt.Printf("  leave: employee %d (%s to %s, %s)\n", l.EmployeeID, l.StartDate, l.EndDate, l.Status)
	}

	// ---- Department Budgets (current month) ----
	fmt.Println("\nSeeding budgets...")
	period := time.Now().Format("2006-01")
	budgets := []models.DepartmentBudget{
		{DepartmentID: deptMap["ENG"], Period: period, Allocated: 50_000_000},
		{DepartmentID: deptMap["MKT"], Period: period, Allocated: 25_000_000},
		{DepartmentID: deptMap["HR"], Period: period, Allocated: 20_000_000},
	}
	for _, b := range budgets {
		if err := db.Create(&b).Error; err != nil {
			log.Fatalf("failed to create budget: %v", err)
		}
		fmt.Printf("  budget: dept_id=%d period=%s allocated=Rp%.0f\n", b.DepartmentID, b.Period, b.Allocated)
	}

	fmt.Println("\nSeed complete!")
}
