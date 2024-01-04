package mapper

import (
	"fmt"
	"github.com/okyirmawan/employee-attendance-api/domain"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/utils"
)

func CheckInRequestToDomain(dto dto.AttendanceRequest) domain.Attendance {
	return domain.Attendance{
		EmployeeID: dto.EmployeeId,
		CheckIn:    utils.CurrentLocalTime(),
		CheckOut:   utils.CurrentLocalTime(),
		LocationIn: fmt.Sprintf("%s,%s", dto.Latitude, dto.Longitude),
		CreatedAt:  utils.CurrentLocalTime(),
		UpdatedAt:  utils.CurrentLocalTime(),
	}
}

func CheckOutRequestToDomain(dto dto.AttendanceRequest) domain.Attendance {
	return domain.Attendance{
		EmployeeID:  dto.EmployeeId,
		CheckOut:    utils.CurrentLocalTime(),
		LocationOut: fmt.Sprintf("%s,%s", dto.Latitude, dto.Longitude),
		UpdatedAt:   utils.CurrentLocalTime(),
	}
}

func DomainToAttendanceResponse(attendance domain.Attendance) dto.AttendanceResponse {
	return dto.AttendanceResponse{
		Id:          attendance.ID,
		EmployeeId:  attendance.EmployeeID,
		CheckIn:     utils.ParseDateTimeToString(attendance.CheckIn),
		CheckOut:    utils.ParseDateTimeToString(attendance.CheckOut),
		LocationIn:  attendance.LocationIn,
		LocationOut: attendance.LocationOut,
	}
}

func DomainToAttendanceResponseList(attendances []domain.Attendance) []dto.AttendanceResponse {
	attendancesDto := make([]dto.AttendanceResponse, len(attendances))

	for i, item := range attendances {
		attendancesDto[i] = DomainToAttendanceResponse(item)
	}

	return attendancesDto
}
