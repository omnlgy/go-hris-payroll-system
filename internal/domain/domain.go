package domain

import (
	"time"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
)

type DepartmentRepository interface {
	Create(department *models.Department) (models.Department, error)
	GetAll() ([]models.Department, error)
	GetByID(id uint) (models.Department, error)
	Update(department *models.Department) (models.Department, error)
	Delete(id uint) error
}

type PositionRepository interface {
	Create(position *models.Position) (models.Position, error)
	GetAll() ([]models.Position, error)
	GetByID(id uint) (models.Position, error)
	Update(position *models.Position) (models.Position, error)
	Delete(id uint) error
}

type EmployeeRepository interface {
	Create(employee *models.Employee) (models.Employee, error)
	GetAll(filter models.FilterEmployee) ([]models.Employee, error)
	GetByID(id uint) (models.Employee, error)
	Update(employee *models.Employee) (models.Employee, error)
	Delete(id uint) error
}

type BlacklistedTokenRepository interface {
	Blacklist(token string, expiredAt time.Time) (*models.BlacklistedToken, error)
	IsBlacklisted(token string) (bool, error)
	CleanupExpired() error
}

type DepartmentBudgetRepository interface {
	Upsert(budget *models.DepartmentBudget) (*models.DepartmentBudget, error)
	GetByDeptAndPeriod(departmentID uint, period string) (*models.DepartmentBudget, error)
	GetAllByPeriod(period string) ([]models.DepartmentBudget, error)
}

type AttendanceRepository interface {
	Create(attendance *models.Attendance) (models.Attendance, error)
	GetAll() ([]models.Attendance, error)
	GetByID(id uint) (models.Attendance, error)
	GetTodayAttendanceEmployee(employeeID uint) (models.Attendance, error)
	GetByEmployeeIDPeriod(employeeID uint, period string) ([]models.Attendance, error)
	Update(attendance *models.Attendance) (models.Attendance, error)
	Delete(id uint) error
}

type LeaveRepository interface {
	Create(leave *models.Leave) (models.Leave, error)
	GetAll() ([]models.Leave, error)
	GetByID(id uint) (models.Leave, error)
	Update(leave *models.Leave) (models.Leave, error)
	Delete(id uint) error
}

type SalaryRepository interface {
	Create(salary *models.Salary) (models.Salary, error)
	GetAll() ([]models.Salary, error)
	GetByID(id uint) (models.Salary, error)
	GetSaleryEmployeeByPeriod(employeeID uint, period string) ([]models.Salary, error)
	Update(salary *models.Salary) (models.Salary, error)
	Delete(id uint) error
}

type AttendanceService interface {
	Attend(employeeID uint) error
}

type DepartmentService interface {
	GetDepartments() ([]models.Department, error)
	GetDepartmentByID(id uint) (models.Department, error)
	CreateDepartment(department *models.Department) (models.Department, error)
	UpdateDepartment(department *models.Department) (models.Department, error)
	DeleteDepartment(id uint) error
}

type PositionService interface {
	GetPositions() ([]models.Position, error)
	GetPositionByID(id uint) (models.Position, error)
	CreatePosition(position *models.Position) (models.Position, error)
	UpdatePosition(position *models.Position) (models.Position, error)
	DeletePosition(id uint) error
}

type EmployeeService interface {
	Add(employee *models.Employee) (models.Employee, error)
	GetEmployees(filter models.FilterEmployee) ([]models.Employee, error)
	Update(employee *models.Employee) (models.Employee, error)
	DeleteEmployee(id uint) error
}

type LeaveService interface {
	RequestLeave(employeeID uint, startDate, endDate, reason string) (models.Leave, error)
	ApproveLeave(leaveID uint, status string) (models.Leave, error)
}

type SalaryService interface {
	GetSaleryEmployeeByPeriod(employeeID uint, period string) ([]models.Salary, error)
	CalculateSalary(employeeID uint, period string) (models.Salary, error)
}
