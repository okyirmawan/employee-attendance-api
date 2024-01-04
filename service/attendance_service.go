package service

import (
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/mapper"
	"github.com/okyirmawan/employee-attendance-api/repository"
)

type AttendanceRepositoryContract interface {
	AttendanceHistory(employeeId uint64) []dto.AttendanceResponse
	CheckIn(dto dto.AttendanceRequest) (dto.AttendanceResponse, error)
	CheckOut(dto dto.AttendanceRequest, id uint64) (dto.AttendanceResponse, error)
}

type AttendanceService struct {
	AttendanceRepository repository.AttendanceRepository
}

func ProviderAttendanceService(m repository.AttendanceRepository) AttendanceService {
	return AttendanceService{
		AttendanceRepository: m,
	}
}

func (m *AttendanceService) CheckIn(dto dto.AttendanceRequest) (dto.AttendanceResponse, error) {
	attendanceEntity := mapper.CheckInRequestToDomain(dto)
	attendance, err := m.AttendanceRepository.CheckIn(attendanceEntity)

	return mapper.DomainToAttendanceResponse(attendance), err
}

func (m *AttendanceService) CheckOut(dto dto.AttendanceRequest, id uint64) (dto.AttendanceResponse, error) {
	attendanceEntity := mapper.CheckOutRequestToDomain(dto)
	attendance, err := m.AttendanceRepository.CheckOut(attendanceEntity, id)

	return mapper.DomainToAttendanceResponse(attendance), err
}

func (m *AttendanceService) AttendanceHistory(employeeId uint64) []dto.AttendanceResponse {
	attendances := m.AttendanceRepository.FindByEmployeeId(employeeId)

	return mapper.DomainToAttendanceResponseList(attendances)
}
