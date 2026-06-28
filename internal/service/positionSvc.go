package service

import (
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
)

type PositionService struct {
	repo domain.PositionRepository
}

func NewPositionService(repo domain.PositionRepository) *PositionService {
	return &PositionService{
		repo: repo,
	}
}

func (s *PositionService) GetPositions() ([]models.Position, error) {
	return s.repo.GetAll()
}

func (s *PositionService) GetPositionByID(id uint) (models.Position, error) {
	return s.repo.GetByID(id)
}

func (s *PositionService) CreatePosition(position *models.Position) (models.Position, error) {
	createdPosition, err := s.repo.Create(position)

	if err != nil {
		return models.Position{}, err
	}

	return createdPosition, nil
}

func (s *PositionService) UpdatePosition(position *models.Position) (models.Position, error) {
	updatedPosition, err := s.repo.Update(position)

	if err != nil {
		return models.Position{}, err
	}

	return updatedPosition, nil
}

func (s *PositionService) DeletePosition(id uint) error {
	return s.repo.Delete(id)
}
