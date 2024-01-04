package mapper

import (
	"fmt"
	"github.com/okyirmawan/employee-attendance-api/domain"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/utils"
)

func CheckInRequestToDomain(dto dto.AttendanceRequest, employeeId uint64) domain.Attendance {
	return domain.Attendance{
		EmployeeID: employeeId,
		CheckIn:    utils.CurrentLocalTime(),
		CheckOut:   utils.CurrentLocalTime(),
		LocationIn: fmt.Sprintf("%s,%s", dto.Latitude, dto.Longitude),
		CreatedAt:  utils.CurrentLocalTime(),
		UpdatedAt:  utils.CurrentLocalTime(),
	}
}

func CheckOutRequestToDomain(dto dto.AttendanceRequest, employeeId uint64) domain.Attendance {
	return domain.Attendance{
		EmployeeID:  employeeId,
		CheckOut:    utils.CurrentLocalTime(),
		LocationOut: fmt.Sprintf("%s,%s", dto.Latitude, dto.Longitude),
		UpdatedAt:   utils.CurrentLocalTime(),
	}
}

func DomainToAttendanceResponse(attendance domain.Attendance) dto.AttendanceResponse {
	checkIn, _ := utils.ConvertTimeToLocalZoneString(attendance.CheckIn)
	checkOut, _ := utils.ConvertTimeToLocalZoneString(attendance.CheckOut)

	return dto.AttendanceResponse{
		Id:          attendance.ID,
		EmployeeId:  attendance.EmployeeID,
		CheckIn:     checkIn,
		CheckOut:    checkOut,
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
