package repository

import (
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

var EmployeeNotFound = gorm.ErrRecordNotFound

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) Create(employee *models.Employee) (models.Employee, error) {
	result := r.db.Create(employee)

	return *employee, result.Error
}

func (r *EmployeeRepository) GetByID(id uint) (models.Employee, error) {
	var employee models.Employee
	result := r.db.Preload("Department").Preload("Position").First(&employee, id)
	return employee, result.Error
}

func (r *EmployeeRepository) GetAll(filter models.FilterEmployee) ([]models.Employee, error) {
	var employees []models.Employee
	query := r.db
	if filter.Name != "" {
		query = query.Where("full_name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.DepartmentID != 0 {
		query = query.Where("department_id = ?", filter.DepartmentID)
	}
	if filter.PositionID != 0 {
		query = query.Where("position_id = ?", filter.PositionID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	result := query.Joins("Department").Joins("Position").Find(&employees)
	return employees, result.Error
}

func (r *EmployeeRepository) Update(employee *models.Employee) (models.Employee, error) {
	result := r.db.Save(employee)
	if result.RowsAffected == 0 {
		return models.Employee{}, EmployeeNotFound
	}
	return *employee, result.Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Employee{}, id)
	if result.RowsAffected == 0 {
		return EmployeeNotFound
	}
	return result.Error
}
