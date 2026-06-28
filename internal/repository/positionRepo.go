package repository

import (
	"errors"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

var PositionNotFound = errors.New("Position not found")

type positionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *positionRepository {
	return &positionRepository{
		db: db,
	}
}

func (r *positionRepository) Create(position *models.Position) (models.Position, error) {
	return *position, r.db.Create(position).Error
}

func (r *positionRepository) GetAll() ([]models.Position, error) {
	var positions []models.Position
	return positions, r.db.Find(&positions).Error
}

func (r *positionRepository) GetByID(id uint) (models.Position, error) {
	var position models.Position
	err := r.db.First(&position, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Position{}, PositionNotFound
	}
	return position, err
}

func (r *positionRepository) Update(position *models.Position) (models.Position, error) {
	if _, err := r.GetByID(position.ID); err != nil {
		return models.Position{}, err
	}
	return *position, r.db.Save(position).Error
}

func (r *positionRepository) Delete(id uint) error {
	if _, err := r.GetByID(id); err != nil {
		return err
	}
	return r.db.Delete(&models.Position{}, id).Error
}
