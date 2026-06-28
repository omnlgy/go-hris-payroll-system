package service

import (
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/models"
)

type LeaveService struct {
	repo domain.LeaveRepository
}

func NewLeaveService(repo domain.LeaveRepository) *LeaveService {
	return &LeaveService{
		repo: repo,
	}
}

func (s *LeaveService) RequestLeave(employeeID uint, startDate, endDate, reason string) (models.Leave, error) {
	leave, err := s.repo.Create(&models.Leave{
		EmployeeID: employeeID,
		StartDate:  startDate,
		EndDate:    endDate,
		Reason:     reason,
		Status:     "PENDING",
	})
	return leave, err
}

func (s *LeaveService) ApproveLeave(leaveID uint, status string) (models.Leave, error) {
	leave, err := s.repo.Update(&models.Leave{
		ID:     leaveID,
		Status: status,
	})
	return leave, err
}
