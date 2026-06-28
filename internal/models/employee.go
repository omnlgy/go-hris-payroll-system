package models

import "time"

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	NIK          string     `gorm:"uniqueIndex;not null;size:20" json:"nik" validate:"required"`
	FullName     string     `gorm:"not null;size:100" json:"full_name" validate:"required"`
	Email        string     `gorm:"uniqueIndex;not null;size:100" json:"email" validate:"required,email"`
	DepartmentID uint       `gorm:"not null" json:"department_id" validate:"required"`
	PositionID   uint       `gorm:"not null" json:"position_id" validate:"required"`
	Role         string     `gorm:"not null;size:20;default:EMPLOYEE" json:"role" validate:"required,oneof=EMPLOYEE HRD ADMIN"`
	Password     string     `gorm:"not null;size:255" json:"-" validate:"required,min=6"`
	Status       string     `gorm:"not null;size:20;default:ACTIVE" json:"status" validate:"required,oneof=ACTIVE SUSPENDED TERMINATED"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"`
	Position     Position   `gorm:"foreignKey:PositionID" json:"position"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (Employee) TableName() string {
	return "employees"
}
