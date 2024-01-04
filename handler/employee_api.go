package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/service"
	"net/http"
	"strconv"
)

type EmployeeAPI struct {
	EmployeeService service.EmployeeService
}

func ProviderEmployeeAPI(k service.EmployeeService) EmployeeAPI {
	return EmployeeAPI{EmployeeService: k}
}

// @Summary Create a new employee
// @Description Create a new employee with the provided data.
// @Accept json
// @Produce json
// @Param employee body dto.EmployeeDTO true "Employee data"
// @Success 200 {object} dto.EmployeeDTO
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees [post]
func (m *EmployeeAPI) Create(e echo.Context) error {
	var newDto dto.EmployeeDTO
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
// @Produce json
// @Success 200 {array} dto.EmployeeDTO
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
// @Produce json
// @Param nip path string true "NIP of the employee"
// @Success 200 {object} dto.EmployeeDTO
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees/{nip} [get]
func (m *EmployeeAPI) FindByNip(e echo.Context) error {
	nip := e.Param("nip")
	employee := m.EmployeeService.FindByNip(nip)

	return SuccessResponse(e, http.StatusOK, employee)
}

// @Summary Update Employee
// @Description Update employee details by ID
// @Produce json
// @Param id path integer true "ID of the employee"
// @Param body body dto.EmployeeDTO true "Employee details to be updated"
// @Success 200 {object} dto.EmployeeDTO
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees/{id} [put]
func (m *EmployeeAPI) Update(e echo.Context) error {
	var newDto dto.EmployeeDTO
	if err := e.Bind(&newDto); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, "Invalid body request")
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid employee ID")
	}

	res, err := m.EmployeeService.Update(newDto, id)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

// @Summary Delete Employee
// @Description Delete employee by ID
// @Produce json
// @Param id path integer true "ID of the employee"
// @Success 200 {object} dto.EmployeeDTO
// @Failure 400 {object} ErrorResp
// @Failure 500 {object} ErrorResp
// @Router /employees/{id} [delete]
func (m *EmployeeAPI) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid employee ID")
	}

	err = m.EmployeeService.Delete(uint64(id))
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, nil)
}
