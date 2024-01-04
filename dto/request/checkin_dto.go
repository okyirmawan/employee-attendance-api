package dto_request

type CheckInDTO struct {
	EmployeeId uint64 `json:"employee_id" validate:"required"`
	LocationIn string `json:"location_in" validate:"required"`
}
