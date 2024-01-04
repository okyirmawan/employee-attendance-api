package handler

import (
	"github.com/labstack/echo/v4"
	dto_request "github.com/okyirmawan/employee-attendance-api/dto/request"
	"github.com/okyirmawan/employee-attendance-api/service"
	"net/http"
	"strconv"
)

type AttendanceAPI struct {
	AttendanceService service.AttendanceService
}

func ProviderAttendanceAPI(attendanceService service.AttendanceService) AttendanceAPI {
	return AttendanceAPI{
		AttendanceService: attendanceService,
	}
}

// CheckIn handles the check-in endpoint.
// @Summary Check-in for attendance
// @Description Record check-in for attendance
// @Accept json
// @Produce json
// @Param body body dto_request.CheckInDTO true "Check In"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /attendance/checkin [post]
func (m *AttendanceAPI) CheckIn(e echo.Context) error {
	var newDto dto_request.CheckInDTO
	if err := e.Bind(&newDto); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	res, err := m.AttendanceService.CheckIn(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// CheckOut handles the check-out endpoint.
// @Summary Check-out for attendance
// @Description Record check-out for attendance
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Param body body dto_request.CheckOutDTO true "Check Out"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /attendance/checkout/{id} [put]
func (m *AttendanceAPI) CheckOut(e echo.Context) error {
	var newDto dto_request.CheckOutDTO
	if err := e.Bind(&newDto); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid employee ID")
	}

	res, err := m.AttendanceService.CheckOut(newDto, uint64(id))
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// @Summary Find Attendance by Employee ID
// @Description Retrieve a list of all attendance.
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /attendance/history/{employee_id} [put]
func (m *AttendanceAPI) History(e echo.Context) error {
	employeeId, _ := strconv.Atoi(e.Param("employee_id"))
	attendances := m.AttendanceService.AttendanceHistory(uint64(employeeId))

	return SuccessResponse(e, http.StatusOK, attendances)
}
