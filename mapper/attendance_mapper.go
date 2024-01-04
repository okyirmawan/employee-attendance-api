package mapper

import (
	"github.com/okyirmawan/employee-attendance-api/domain"
	"github.com/okyirmawan/employee-attendance-api/dto"
	dto_request "github.com/okyirmawan/employee-attendance-api/dto/request"
	"github.com/okyirmawan/employee-attendance-api/utils"
)

func CheckInDtoToAttendanceDomain(dto dto_request.CheckInDTO) domain.Attendance {
	return domain.Attendance{
		EmployeeID: dto.EmployeeId,
		CheckIn:    utils.CurrentLocalTime(),
		CheckOut:   utils.CurrentLocalTime(),
		LocationIn: dto.LocationIn,
		CreatedAt:  utils.CurrentLocalTime(),
		UpdatedAt:  utils.CurrentLocalTime(),
	}
}

func CheckOutDtoToAttendanceDomain(dto dto_request.CheckOutDTO) domain.Attendance {
	return domain.Attendance{
		EmployeeID:  dto.EmployeeId,
		CheckOut:    utils.CurrentLocalTime(),
		LocationOut: dto.LocationOut,
		UpdatedAt:   utils.CurrentLocalTime(),
	}
}

func ToAttendanceDomain(dto dto.AttendanceDTO) domain.Attendance {
	return domain.Attendance{
		EmployeeID:  dto.EmployeeId,
		CheckIn:     utils.ParseStringToDateTime(dto.CheckIn),
		CheckOut:    utils.ParseStringToDateTime(dto.CheckOut),
		TotalHours:  0,
		LocationIn:  "",
		LocationOut: "",
		CreatedAt:   utils.CurrentLocalTime(),
		UpdatedAt:   utils.CurrentLocalTime(),
	}
}

func ToAttendanceDto(attendance domain.Attendance) dto.AttendanceDTO {
	return dto.AttendanceDTO{
		Id:          attendance.ID,
		EmployeeId:  attendance.EmployeeID,
		CheckIn:     utils.ParseDateTimeToString(attendance.CheckIn),
		CheckOut:    utils.ParseDateTimeToString(attendance.CheckOut),
		LocationIn:  attendance.LocationIn,
		LocationOut: attendance.LocationOut,
	}
}

func ToAttendanceDtoList(attendances []domain.Attendance) []dto.AttendanceDTO {
	dtos := make([]dto.AttendanceDTO, len(attendances))

	for i, item := range attendances {
		dtos[i] = ToAttendanceDto(item)
	}

	return dtos
}
