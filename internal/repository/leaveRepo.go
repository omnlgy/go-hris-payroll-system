package repository

import (
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

type LeaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) *LeaveRepository {
	return &LeaveRepository{
		db: db,
	}
}

func (r *LeaveRepository) Create(leave *models.Leave) (models.Leave, error) {
	if err := r.db.Create(leave).Error; err != nil {
		return models.Leave{}, err
	}
	return *leave, r.db.Preload("Employee").First(leave, leave.ID).Error
}

func (r *LeaveRepository) GetAll() ([]models.Leave, error) {
	return []models.Leave{}, nil
}

func (r *LeaveRepository) GetByID(id uint) (models.Leave, error) {
	return models.Leave{}, nil
}

func (r *LeaveRepository) Update(leave *models.Leave) (models.Leave, error) {
	if err := r.db.Model(&models.Leave{}).Where("id = ?", leave.ID).Updates(map[string]interface{}{
		"status": leave.Status,
	}).Error; err != nil {
		return models.Leave{}, err
	}
	var result models.Leave
	err := r.db.Preload("Employee").First(&result, leave.ID).Error
	return result, err
}

func (r *LeaveRepository) Delete(id uint) error {
	return nil
}
