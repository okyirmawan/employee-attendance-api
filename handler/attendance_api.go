package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/service"
	"github.com/okyirmawan/employee-attendance-api/utils/token"
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
// @Tags Attendances
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param body body dto.AttendanceRequest true "Attendance Request"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /attendances/checkin [post]
func (m *AttendanceAPI) CheckIn(e echo.Context) error {
	employeeId, err := token.ExtractTokenID(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var request dto.AttendanceRequest
	if err := e.Bind(&request); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	res, err := m.AttendanceService.CheckIn(request, employeeId)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// CheckOut handles the check-out endpoint.
// @Summary Check-out for attendance
// @Description Record check-out for attendance
// @Tags Attendances
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path int true "Attendance ID"
// @Param body body dto.AttendanceRequest true "Check Out"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /attendances/checkout/{id} [put]
func (m *AttendanceAPI) CheckOut(e echo.Context) error {
	employeeId, err := token.ExtractTokenID(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var request dto.AttendanceRequest
	if err := e.Bind(&request); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	attendanceId, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid attendance ID")
	}

	res, err := m.AttendanceService.CheckOut(request, employeeId, uint64(attendanceId))
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// @Summary Find Attendance by Employee ID
// @Description Retrieve a list of all attendance.
// @Tags Attendances
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /attendances/history [get]
func (m *AttendanceAPI) History(e echo.Context) error {
	employeeId, err := token.ExtractTokenID(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	attendances := m.AttendanceService.AttendanceHistory(employeeId)

	return SuccessResponse(e, http.StatusOK, attendances)
}
