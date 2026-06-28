package repository

import (
	"errors"

	"github.com/omnlgy/go-hris-payroll-system/internal/dto"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

var DepartmentNotFound = errors.New("Department not found")

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) GetAll() ([]dto.DepartmentRepository, error) {
	var depts []models.Department
	if err := r.db.Find(&depts).Error; err != nil {
		return nil, err
	}
	return dto.DepartmentFromModels(depts), nil
}

func (r *departmentRepository) GetByID(id uint) (dto.DepartmentRepository, error) {
	var dept models.Department
	err := r.db.First(&dept, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.DepartmentRepository{}, DepartmentNotFound
	}
	if err != nil {
		return dto.DepartmentRepository{}, err
	}
	return dto.DepartmentFromModel(dept), nil
}

func (r *departmentRepository) Create(input dto.DepartmentCreate) (dto.DepartmentRepository, error) {
	dept := input.ToModel()
	if err := r.db.Create(&dept).Error; err != nil {
		return dto.DepartmentRepository{}, err
	}
	return dto.DepartmentFromModel(dept), nil
}

func (r *departmentRepository) Update(input dto.DepartmentUpdate) (dto.DepartmentRepository, error) {
	dept := input.ToModel()
	result := r.db.Save(&dept)
	if result.RowsAffected == 0 {
		return dto.DepartmentRepository{}, DepartmentNotFound
	}
	if result.Error != nil {
		return dto.DepartmentRepository{}, result.Error
	}
	return dto.DepartmentFromModel(dept), nil
}

func (r *departmentRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Department{}, id)

	if result.RowsAffected == 0 {
		return DepartmentNotFound
	}

	return result.Error
}
