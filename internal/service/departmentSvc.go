package service

import (
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
)

type DepartmentService struct {
	repo domain.DepartmentRepository
}

func NewDepartmentService(repo domain.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		repo: repo,
	}
}

func (s *DepartmentService) GetDepartments() ([]dto.DepartmentRepository, error) {
	return s.repo.GetAll()
}

func (s *DepartmentService) GetDepartmentByID(id uint) (dto.DepartmentRepository, error) {
	return s.repo.GetByID(id)
}

func (s *DepartmentService) CreateDepartment(input dto.DepartmentCreate) (dto.DepartmentRepository, error) {
	return s.repo.Create(input)
}

func (s *DepartmentService) UpdateDepartment(input dto.DepartmentUpdate) (dto.DepartmentRepository, error) {
	return s.repo.Update(input)
}

func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.repo.Delete(id)
}
