package dto

import (
	"time"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
)

// DepartmentRepository is the repository-layer data transfer object.
// Decouples the service layer from GORM model internals.
type DepartmentRepository struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DepartmentCreate is used when creating a department.
// Excludes auto-generated fields (ID, timestamps).
type DepartmentCreate struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

// DepartmentUpdate is used when updating a department.
type DepartmentUpdate struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

// ToModel converts a create DTO to a GORM model.
func (d DepartmentCreate) ToModel() models.Department {
	return models.Department{
		Name: d.Name,
		Code: d.Code,
	}
}

// ToModel converts an update DTO to a GORM model.
func (d DepartmentUpdate) ToModel() models.Department {
	return models.Department{
		ID:   d.ID,
		Name: d.Name,
		Code: d.Code,
	}
}

// FromModel converts a GORM model to a repository DTO.
func DepartmentFromModel(m models.Department) DepartmentRepository {
	return DepartmentRepository{
		ID:        m.ID,
		Name:      m.Name,
		Code:      m.Code,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

// DepartmentFromModels converts a slice of GORM models to DTOs.
func DepartmentFromModels(models []models.Department) []DepartmentRepository {
	dtos := make([]DepartmentRepository, len(models))
	for i, m := range models {
		dtos[i] = DepartmentFromModel(m)
	}
	return dtos
}
