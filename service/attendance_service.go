package service

import (
	"fmt"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/mapper"
	"github.com/okyirmawan/employee-attendance-api/repository"
	"time"
)

type AttendanceRepositoryContract interface {
	AttendanceHistory(employeeId uint64) []dto.AttendanceResponse
	CheckIn(dto dto.AttendanceRequest, employeeId uint64) (dto.AttendanceResponse, error)
	CheckOut(dto dto.AttendanceRequest, employeeId uint64, attendanceId uint64) (dto.AttendanceResponse, error)
}

type AttendanceService struct {
	AttendanceRepository repository.AttendanceRepository
}

func ProviderAttendanceService(m repository.AttendanceRepository) AttendanceService {
	return AttendanceService{
		AttendanceRepository: m,
	}
}

func (m *AttendanceService) CheckIn(request dto.AttendanceRequest, employeeId uint64) (dto.AttendanceResponse, error) {
	attendance := m.AttendanceRepository.FindByDateAndEmployeeId(time.Now().Format("2006-01-02"), employeeId)
	if attendance.ID > 0 {
		return dto.AttendanceResponse{}, fmt.Errorf("employee already check in")
	}

	attendanceEntity := mapper.CheckInRequestToDomain(request, employeeId)
	attendance, err := m.AttendanceRepository.CheckIn(attendanceEntity)

	return mapper.DomainToAttendanceResponse(attendance), err
}

func (m *AttendanceService) CheckOut(dto dto.AttendanceRequest, employeeId uint64, attendanceid uint64) (dto.AttendanceResponse, error) {
	attendanceEntity := mapper.CheckOutRequestToDomain(dto, employeeId)
	attendance, err := m.AttendanceRepository.CheckOut(attendanceEntity, attendanceid)

	return mapper.DomainToAttendanceResponse(attendance), err
}

func (m *AttendanceService) AttendanceHistory(employeeId uint64) []dto.AttendanceResponse {
	attendances := m.AttendanceRepository.FindByEmployeeId(employeeId)

	return mapper.DomainToAttendanceResponseList(attendances)
}
