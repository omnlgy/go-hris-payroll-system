package models

import "time"

type Leave struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `gorm:"not null;index" json:"employee_id" validate:"required"`
	StartDate  string    `gorm:"not null;type:date;size:10" json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate    string    `gorm:"not null;type:date;size:10" json:"end_date" validate:"required,datetime=2006-01-02"`
	Reason     string    `gorm:"not null;type:text" json:"reason" validate:"required"`
	Status     string    `gorm:"not null;size:10;default:PENDING" json:"status" validate:"required,oneof=PENDING APPROVED REJECTED"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Leave) TableName() string {
	return "leaves"
}
