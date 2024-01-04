package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/service"
	"github.com/okyirmawan/employee-attendance-api/utils/token"
	"net/http"
)

type EmployeeAPI struct {
	EmployeeService service.EmployeeService
}

func ProviderEmployeeAPI(k service.EmployeeService) EmployeeAPI {
	return EmployeeAPI{EmployeeService: k}
}

// @Summary Create a new employee
// @Description Create a new employee with the provided data.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body dto.EmployeeRequest true "Employee data"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees [post]
func (m *EmployeeAPI) Create(e echo.Context) error {
	var newDto dto.EmployeeRequest
	if err := e.Bind(&newDto); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	res, err := m.EmployeeService.Create(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// @Summary Retrieve all employees
// @Description Retrieve a list of all employees.
// @Tags Employees
// @Produce json
// @Success 200 {array} SuccessResp "Success response"
// @Failure 204 {object} ErrorResp
// @Router /employees [get]
func (m *EmployeeAPI) FindAll(e echo.Context) error {

	employees := m.EmployeeService.FindAll()

	if len(employees) == 0 {
		return SuccessResponse(e, http.StatusNoContent, employees)
	}

	return SuccessResponse(e, http.StatusOK, employees)
}

// @Summary Find Employee by NIP
// @Description Get employee details by NIP
// @Tags Employees
// @Produce json
// @Param nip path string true "NIP of the employee"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees/{nip} [get]
func (m *EmployeeAPI) FindByNip(e echo.Context) error {
	nip := e.Param("nip")
	employee := m.EmployeeService.FindByNip(nip)

	return SuccessResponse(e, http.StatusOK, employee)
}

// @Summary Update Employee
// @Description Update employee details by authentication token
// @Tags Employees
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param body body dto.EmployeeRequest true "Employee details to be updated"
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees [put]
func (m *EmployeeAPI) Update(e echo.Context) error {
	var newDto dto.EmployeeRequest
	if err := e.Bind(&newDto); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	employeeId, err := token.ExtractTokenID(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res, err := m.EmployeeService.Update(newDto, employeeId)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// @Summary Delete Employee
// @Description Delete employee by authentication token
// @Tags Employees
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees [delete]
func (m *EmployeeAPI) Delete(e echo.Context) error {
	employeeId, err := token.ExtractTokenID(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = m.EmployeeService.Delete(employeeId)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, nil)
}

// @Summary Login Employee
// @Description Login with username and password to get token.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body dto.LoginRequest true "Login"
// @Success 200 {string} SuccessResp "Success response"
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees/login [post]
func (m *EmployeeAPI) Login(e echo.Context) error {
	var request dto.LoginRequest
	if err := e.Bind(&request); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	auth, err := m.EmployeeService.Auth(request.Username, request.Password)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, auth)
}
