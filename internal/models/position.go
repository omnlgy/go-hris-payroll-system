package models

import "time"

type Position struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null;size:100" json:"title" validate:"required"`
	BaseSalary float64   `gorm:"not null" json:"base_salary" validate:"required,gt=0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Position) TableName() string {
	return "positions"
}
