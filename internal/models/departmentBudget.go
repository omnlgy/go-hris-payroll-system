package models

import "time"

type DepartmentBudget struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	DepartmentID uint      `gorm:"not null;uniqueIndex:idx_dept_period" json:"department_id" validate:"required"`
	Period       string    `gorm:"not null;size:7;uniqueIndex:idx_dept_period" json:"period" validate:"required,datetime=2006-01"`
	Allocated    float64   `gorm:"not null" json:"allocated" validate:"required,gt=0"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (DepartmentBudget) TableName() string {
	return "department_budgets"
}