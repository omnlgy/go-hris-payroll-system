package repository

import (
	"time"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{
		db: db,
	}
}

func (r *AttendanceRepository) Create(attendance *models.Attendance) (models.Attendance, error) {
	return *attendance, r.db.Create(attendance).Error
}

func (r *AttendanceRepository) GetByID(id uint) (models.Attendance, error) {
	var attendance models.Attendance
	return attendance, r.db.First(&attendance, id).Error
}

func (r *AttendanceRepository) GetTodayAttendanceEmployee(employeeID uint) (models.Attendance, error) {
	var attendance models.Attendance
	return attendance, r.db.Where("date = ? AND employee_id = ?", time.Now().Format("2006-01-02"), employeeID).First(&attendance).Error
}

func (r *AttendanceRepository) GetAll() ([]models.Attendance, error) {
	var attendances []models.Attendance
	return attendances, r.db.Find(&attendances).Error
}

func (r *AttendanceRepository) GetByEmployeeIDPeriod(employeeID uint, period string) ([]models.Attendance, error) {
	var attendances []models.Attendance
	return attendances, r.db.Where("employee_id = ? AND date LIKE ?", employeeID, period+"%").Find(&attendances).Error
}

func (r *AttendanceRepository) Update(attendance *models.Attendance) (models.Attendance, error) {
	return *attendance, r.db.Save(attendance).Error
}

func (r *AttendanceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Attendance{}, id).Error
}
