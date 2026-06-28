package service

import (
	"errors"

	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService struct {
	repo      domain.EmployeeRepository
	deptSvc   domain.DepartmentService
	posSvc    domain.PositionService
}

func NewEmployeeService(repo domain.EmployeeRepository, deptSvc domain.DepartmentService, posSvc domain.PositionService) *EmployeeService {
	return &EmployeeService{
		repo:      repo,
		deptSvc:   deptSvc,
		posSvc:    posSvc,
	}
}

func (s *EmployeeService) Add(employee *models.Employee) (models.Employee, error) {
	// Validate department and position
	if _, err := s.deptSvc.GetDepartmentByID(employee.DepartmentID); err != nil {
		return models.Employee{}, errors.New("department not found")
	}
	if _, err := s.posSvc.GetPositionByID(employee.PositionID); err != nil {
		return models.Employee{}, errors.New("position not found")
	}

	// Hash password if provided
	if employee.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.Employee{}, err
		}
		employee.Password = string(hashedPassword)
	}

	// Set default role if empty
	if employee.Role == "" {
		employee.Role = "EMPLOYEE"
	}

	return s.repo.Create(employee)
}

func (s *EmployeeService) GetEmployees(filter models.FilterEmployee) ([]models.Employee, error) {
	return s.repo.GetAll(filter)
}

func (s *EmployeeService) GetByID(id uint) (models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) Update(employee *models.Employee) (models.Employee, error) {
	// Validate department and position
	if _, err := s.deptSvc.GetDepartmentByID(employee.DepartmentID); err != nil {
		return models.Employee{}, errors.New("department not found")
	}
	if _, err := s.posSvc.GetPositionByID(employee.PositionID); err != nil {
		return models.Employee{}, errors.New("position not found")
	}

	// Hash password if provided (not empty)
	if employee.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.Employee{}, err
		}
		employee.Password = string(hashedPassword)
	}

	return s.repo.Update(employee)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	return s.repo.Delete(id)
}