package service

import (
	"github.com/okyirmawan/employee-attendance-api/dto"
	dto_request "github.com/okyirmawan/employee-attendance-api/dto/request"
	"github.com/okyirmawan/employee-attendance-api/mapper"
	"github.com/okyirmawan/employee-attendance-api/repository"
)

type AttendanceRepositoryContract interface {
	AttendanceHistory(employeeId uint64) []dto.AttendanceDTO
	CheckIn(dto dto_request.CheckInDTO) (dto.AttendanceDTO, error)
	CheckOut(dto dto_request.CheckOutDTO, id uint64) (dto.AttendanceDTO, error)
}

type AttendanceService struct {
	AttendanceRepository repository.AttendanceRepository
}

func ProviderAttendanceService(m repository.AttendanceRepository) AttendanceService {
	return AttendanceService{
		AttendanceRepository: m,
	}
}

func (m *AttendanceService) CheckIn(dto dto_request.CheckInDTO) (dto.AttendanceDTO, error) {
	attendanceEntity := mapper.CheckInDtoToAttendanceDomain(dto)
	attendance, err := m.AttendanceRepository.CheckIn(attendanceEntity)

	return mapper.ToAttendanceDto(attendance), err
}

func (m *AttendanceService) CheckOut(dto dto_request.CheckOutDTO, id uint64) (dto.AttendanceDTO, error) {
	attendanceEntity := mapper.CheckOutDtoToAttendanceDomain(dto)
	attendance, err := m.AttendanceRepository.CheckOut(attendanceEntity, id)

	return mapper.ToAttendanceDto(attendance), err
}

func (m *AttendanceService) AttendanceHistory(employeeId uint64) []dto.AttendanceDTO {
	attendances := m.AttendanceRepository.FindByEmployeeId(employeeId)

	return mapper.ToAttendanceDtoList(attendances)
}
