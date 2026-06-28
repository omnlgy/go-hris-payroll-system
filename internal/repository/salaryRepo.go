package repository

import (
	"errors"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

var SalaryNotFound = errors.New("Salary not found")

type salaryRepository struct {
	db *gorm.DB
}

func NewSalaryRepository(db *gorm.DB) *salaryRepository {
	return &salaryRepository{db: db}
}

func (r *salaryRepository) Create(salary *models.Salary) (models.Salary, error) {
	if err := r.db.Create(salary).Error; err != nil {
		return models.Salary{}, err
	}

	var result models.Salary
	err := r.db.Preload("Employee").First(&result, salary.ID).Error
	return result, err
}

func (r *salaryRepository) GetAll() ([]models.Salary, error) {
	var salaries []models.Salary
	return salaries, r.db.Preload("Employee").Find(&salaries).Error
}

func (r *salaryRepository) GetByID(id uint) (models.Salary, error) {
	var salary models.Salary
	err := r.db.Preload("Employee").First(&salary, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Salary{}, SalaryNotFound
	}
	return salary, err
}

func (r *salaryRepository) GetSaleryEmployeeByPeriod(employeeID uint, period string) ([]models.Salary, error) {
	var salaries []models.Salary
	return salaries, r.db.Preload("Employee").Where("employee_id = ? AND period = ?", employeeID, period).Find(&salaries).Error
}

func (r *salaryRepository) Update(salary *models.Salary) (models.Salary, error) {
	if _, err := r.GetByID(salary.ID); err != nil {
		return models.Salary{}, err
	}

	if err := r.db.Save(salary).Error; err != nil {
		return models.Salary{}, err
	}

	var result models.Salary
	err := r.db.Preload("Employee").First(&result, salary.ID).Error
	return result, err
}

func (r *salaryRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Salary{}, id)
	if result.RowsAffected == 0 {
		return SalaryNotFound
	}
	return result.Error
}
