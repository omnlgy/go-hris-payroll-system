package repository

import (
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

var DepartmentBudgetNotFound = gorm.ErrRecordNotFound

type DepartmentBudgetRepository struct {
	db *gorm.DB
}

func NewDepartmentBudgetRepository(db *gorm.DB) *DepartmentBudgetRepository {
	return &DepartmentBudgetRepository{
		db: db,
	}
}

func (r *DepartmentBudgetRepository) Upsert(budget *models.DepartmentBudget) (*models.DepartmentBudget, error) {
	var existing models.DepartmentBudget
	err := r.db.Where("department_id = ? AND period = ?", budget.DepartmentID, budget.Period).
		First(&existing).Error

	if err == nil {
		// Update existing
		existing.Allocated = budget.Allocated
		if err := r.db.Save(&existing).Error; err != nil {
			return nil, err
		}
		return &existing, nil
	}

	if err == gorm.ErrRecordNotFound {
		// Create new
		if err := r.db.Create(budget).Error; err != nil {
			return nil, err
		}
		return budget, nil
	}

	return nil, err
}

func (r *DepartmentBudgetRepository) GetByDeptAndPeriod(departmentID uint, period string) (*models.DepartmentBudget, error) {
	var budget models.DepartmentBudget
	err := r.db.Where("department_id = ? AND period = ?", departmentID, period).
		Preload("Department").
		First(&budget).Error
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *DepartmentBudgetRepository) GetAllByPeriod(period string) ([]models.DepartmentBudget, error) {
	var budgets []models.DepartmentBudget
	err := r.db.Where("period = ?", period).
		Preload("Department").
		Find(&budgets).Error
	return budgets, err
}