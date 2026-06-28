package service

import (
	"fmt"

	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
)

type SalaryService struct {
	repo           domain.SalaryRepository
	attendanceRepo domain.AttendanceRepository
	employeeRepo   domain.EmployeeRepository
}

func NewSalaryService(repo domain.SalaryRepository, attendanceRepo domain.AttendanceRepository, employeeRepo domain.EmployeeRepository) *SalaryService {
	return &SalaryService{
		repo:           repo,
		attendanceRepo: attendanceRepo,
		employeeRepo:   employeeRepo,
	}
}

func (s *SalaryService) GetSaleryEmployeeByPeriod(employeeID uint, period string) ([]models.Salary, error) {
	return s.repo.GetSaleryEmployeeByPeriod(employeeID, period)
}

func (s *SalaryService) CalculateSalary(employeeID uint, period string) (models.Salary, error) {
	// 1. Get employee (with position to get base salary)
	employee, err := s.employeeRepo.GetByID(employeeID)
	if err != nil {
		return models.Salary{}, fmt.Errorf("employee not found: %w", err)
	}

	// 2. Get attendances for this period
	attendances, err := s.attendanceRepo.GetByEmployeeIDPeriod(employeeID, period)
	if err != nil {
		return models.Salary{}, fmt.Errorf("failed to get attendances: %w", err)
	}

	// 3. Count PRESENT, LATE, ABSENT days
	var presentDays, lateDays, absentDays int64
	for _, a := range attendances {
		switch a.Status {
		case "PRESENT":
			presentDays++
		case "LATE":
			lateDays++
		case "ABSENT":
			absentDays++
		}
	}

	// 4. Calculate salary components
	//    Allowance: Rp 50.000 per on-time (PRESENT) day
	//    Deductions: Rp 20.000 per late day, Rp 100.000 per absent day
	baseSalary := employee.Position.BaseSalary

	const (
		attendanceBonus = 50_000.0
		latePenalty     = 20_000.0
		absentPenalty   = 100_000.0
	)

	allowance := float64(presentDays) * attendanceBonus
	deductions := float64(lateDays)*latePenalty + float64(absentDays)*absentPenalty

	// 5. Check if salary already exists for this period → update, else create
	existingSalaries, _ := s.repo.GetSaleryEmployeeByPeriod(employeeID, period)
	if len(existingSalaries) > 0 {
		existing := existingSalaries[0]
		existing.BasicSalary = baseSalary
		existing.Allowance = allowance
		existing.Deductions = deductions
		// NetSalary auto-calculated by BeforeSave hook
		return s.repo.Update(&existing)
	}

	// Create new salary record
	salary := &models.Salary{
		EmployeeID:  employeeID,
		Period:      period,
		BasicSalary: baseSalary,
		Allowance:   allowance,
		Deductions:  deductions,
	}

	return s.repo.Create(salary)
}
