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

	// ---- Departments ----
	depts := []models.Department{
		{Name: "Engineering", Code: "ENG"},
		{Name: "Marketing", Code: "MKT"},
		{Name: "Human Resources", Code: "HR"},
	}
	for i := range depts {
		if err := db.FirstOrCreate(&depts[i], models.Department{Code: depts[i].Code}).Error; err != nil {
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
		if err := db.FirstOrCreate(&pos, models.Position{Title: p.Title}).Error; err != nil {
			log.Fatalf("failed to create position %s: %v", p.Title, err)
		}
		positions = append(positions, pos)
		fmt.Printf("  position: %s (Rp%.0f, id=%d)\n", pos.Title, pos.BaseSalary, pos.ID)
	}

	// map positions for easy lookup
	posMap := make(map[string]uint)
	for _, p := range positions {
		posMap[p.Title] = p.ID
	}

	// dept map
	deptMap := make(map[string]uint)
	for _, d := range depts {
		deptMap[d.Code] = d.ID
	}

	// ---- Employees (5) ----
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
		{NIK: "001", FullName: "Budi Santoso", Email: "budi@company.com", DeptCode: "ENG", PosTitle: "Engineering Manager", Role: "ADMIN", Password: "password123"},
		{NIK: "002", FullName: "Siti Rahayu", Email: "siti@company.com", DeptCode: "ENG", PosTitle: "Senior Engineer", Role: "EMPLOYEE", Password: "password123"},
		{NIK: "003", FullName: "Agus Wijaya", Email: "agus@company.com", DeptCode: "ENG", PosTitle: "Junior Engineer", Role: "EMPLOYEE", Password: "password123"},
		{NIK: "004", FullName: "Dewi Lestari", Email: "dewi@company.com", DeptCode: "MKT", PosTitle: "Marketing Manager", Role: "MANAGER", Password: "password123"},
		{NIK: "005", FullName: "Rudi Hartono", Email: "rudi@company.com", DeptCode: "HR", PosTitle: "HR Staff", Role: "EMPLOYEE", Password: "password123"},
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
		if err := db.FirstOrCreate(&emp, models.Employee{NIK: e.NIK}).Error; err != nil {
			log.Fatalf("failed to create employee %s: %v", e.NIK, err)
		}
		fmt.Printf("  employee: %s (%s, %s)\n", emp.FullName, emp.Role, emp.Status)
	}

	// ---- Department Budgets (current month) ----
	period := time.Now().Format("2006-01")
	budgets := []models.DepartmentBudget{
		{DepartmentID: deptMap["ENG"], Period: period, Allocated: 50_000_000},
		{DepartmentID: deptMap["MKT"], Period: period, Allocated: 25_000_000},
		{DepartmentID: deptMap["HR"], Period: period, Allocated: 20_000_000},
	}
	for _, b := range budgets {
		// use upsert-like: delete then create
		var existing models.DepartmentBudget
		err := db.Where("department_id = ? AND period = ?", b.DepartmentID, b.Period).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&b).Error; err != nil {
				log.Fatalf("failed to create budget: %v", err)
			}
			fmt.Printf("  budget: %s period=%s allocated=Rp%.0f\n",
				depts[b.DepartmentID-1].Name, b.Period, b.Allocated)
		} else if err == nil {
			fmt.Printf("  budget already exists for %s period=%s\n",
				depts[b.DepartmentID-1].Name, b.Period)
		} else {
			log.Fatalf("failed to check budget: %v", err)
		}
	}

	fmt.Println("\nSeed complete!")
}