package service

import (
	"errors"
	"time"

	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"gorm.io/gorm"
)

type AttendanceService struct {
	repo domain.AttendanceRepository
}

func NewAttendanceService(repo domain.AttendanceRepository) *AttendanceService {
	return &AttendanceService{
		repo: repo,
	}
}

func (s *AttendanceService) Attend(employeeId uint) error {
	todayAtendance, err := s.repo.GetTodayAttendanceEmployee(employeeId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	isCheckedIn := err == nil

	if isCheckedIn {
		todayAtendance.CheckOut = time.Now().Format("15:04")
		if _, err := s.repo.Update(&todayAtendance); err != nil {
			return err
		}
	} else {
		todayAtendance.EmployeeID = employeeId
		todayAtendance.Date = time.Now().Format("2006-01-02")
		todayAtendance.CheckIn = time.Now().Format("15:04")

		minimumTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 0, 0, 0, time.Local)
		if time.Now().After(minimumTime) {
			todayAtendance.Status = "LATE"
		}

		if _, err := s.repo.Create(&todayAtendance); err != nil {
			return err
		}
	}

	return nil
}
