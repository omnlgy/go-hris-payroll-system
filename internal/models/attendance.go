package models

import "time"

type Attendance struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `gorm:"not null;index" json:"employee_id" validate:"required"`
	Date       string    `gorm:"not null;type:date;size:10" json:"date" validate:"required,datetime=2006-01-02"`
	CheckIn    string    `gorm:"not null;size:5" json:"check_in" validate:"required,datetime=15:04"`
	CheckOut   string    `gorm:"size:5" json:"check_out" validate:"datetime=15:04"`
	Status     string    `gorm:"size:10;default:PRESENT" json:"status" validate:"oneof=PRESENT LATE ABSENT"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Attendance) TableName() string {
	return "attendances"
}
