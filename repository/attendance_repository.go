package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/okyirmawan/employee-attendance-api/domain"
)

type AttendanceRepositoryContract interface {
	FindById(id uint64) (domain.Attendance, error)
	FindByEmployeeId(id uint64) []domain.Attendance
	CheckIn(attendance domain.Attendance) (domain.Attendance, error)
	CheckOut(attendance domain.Attendance, id uint64) (domain.Attendance, error)
}

type AttendanceRepository struct {
	DB *gorm.DB
}

func ProviderAttendanceRepository(DB *gorm.DB) AttendanceRepository {
	return AttendanceRepository{DB: DB}
}

func (m *AttendanceRepository) FindById(id uint64) (domain.Attendance, error) {
	var attendance domain.Attendance

	result := m.DB.Where("id = ?", id).First(&attendance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Record not found
			return attendance, fmt.Errorf("attendance with ID %d not found", id)
		}
		// Other error
		return attendance, result.Error
	}

	return attendance, nil
}

func (m *AttendanceRepository) FindByEmployeeId(id uint64) []domain.Attendance {
	var attendances []domain.Attendance

	m.DB.Where("employee_id =? ", id).Find(&attendances)

	return attendances
}

func (m *AttendanceRepository) CheckIn(attendance domain.Attendance) (domain.Attendance, error) {
	if err := m.DB.Create(&attendance).Error; err != nil {
		return attendance, err
	}

	return attendance, nil
}

//func (m *AttendanceRepository) CheckOut(attendance domain.Attendance, id uint64) (domain.Attendance, error) {
//	attendance.ID = id
//
//	if err := m.DB.Model(&domain.Attendance{}).Where("id = ?", id).Update(attendance).Error; err != nil {
//		return attendance, err
//	}
//
//	return attendance, nil
//}

func (m *AttendanceRepository) CheckOut(attendance domain.Attendance, id uint64) (domain.Attendance, error) {
	// Fetch the existing attendance from the database
	existingAttendance := domain.Attendance{}
	if err := m.DB.First(&existingAttendance, id).Error; err != nil {
		return existingAttendance, err
	}

	// Update the specific field
	existingAttendance.CheckOut = attendance.CheckOut
	existingAttendance.LocationOut = attendance.LocationOut

	// Save the updated attendance back to the database
	if err := m.DB.Save(&existingAttendance).Error; err != nil {
		return existingAttendance, err
	}

	return existingAttendance, nil
}
