package models

import "time"

type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;size:100" json:"name" validate:"required"`
	Code      string    `gorm:"uniqueIndex;not null;size:20" json:"code" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Department) TableName() string {
	return "departments"
}
