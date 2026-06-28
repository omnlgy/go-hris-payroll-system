package models

import (
	"time"

	"gorm.io/gorm"
)

type Salary struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `gorm:"not null;index" json:"employee_id" validate:"required"`
	Period      string    `gorm:"not null;size:7" json:"period" validate:"required,datetime=2006-01"`
	BasicSalary float64   `gorm:"not null" json:"basic_salary" validate:"required,gt=0"`
	Allowance   float64   `gorm:"not null;default:0" json:"allowance"`
	Deductions  float64   `gorm:"not null;default:0" json:"deductions"`
	NetSalary   float64   `gorm:"not null" json:"net_salary"`
	Employee    Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Salary) TableName() string {
	return "salaries"
}

// BeforeSave hook — auto-calculates NetSalary before insert/update.
func (s *Salary) BeforeSave(_ *gorm.DB) (err error) {
	s.NetSalary = s.BasicSalary + s.Allowance - s.Deductions
	return
}
