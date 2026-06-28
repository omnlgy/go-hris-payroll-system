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

func (r *EmployeeRepository) GetEmployeeByEmail(email string) (models.Employee, error) {
	var employee models.Employee
	result := r.db.Preload("Department").Preload("Position").Where("email = ?", email).First(&employee)
	return employee, result.Error
}

func (r *EmployeeRepository) Update(employee *models.Employee) (models.Employee, error) {
	// Use Updates with map to only set non-zero fields (partial update)
	updates := map[string]interface{}{}
	if employee.NIK != "" {
		updates["nik"] = employee.NIK
	}
	if employee.FullName != "" {
		updates["full_name"] = employee.FullName
	}
	if employee.Email != "" {
		updates["email"] = employee.Email
	}
	if employee.DepartmentID != 0 {
		updates["department_id"] = employee.DepartmentID
	}
	if employee.PositionID != 0 {
		updates["position_id"] = employee.PositionID
	}
	if employee.Role != "" {
		updates["role"] = employee.Role
	}
	if employee.Status != "" {
		updates["status"] = employee.Status
	}
	if employee.Password != "" {
		updates["password"] = employee.Password
	}

	if len(updates) == 0 {
		return *employee, nil
	}

	if err := r.db.Model(&models.Employee{}).Where("id = ?", employee.ID).Updates(updates).Error; err != nil {
		return models.Employee{}, err
	}
	var result models.Employee
	err := r.db.Preload("Department").Preload("Position").First(&result, employee.ID).Error
	return result, err
}

func (r *EmployeeRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Employee{}, id)
	if result.RowsAffected == 0 {
		return EmployeeNotFound
	}
	return result.Error
}
