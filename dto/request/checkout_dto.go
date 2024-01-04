package dto_request

type CheckOutDTO struct {
	EmployeeId  uint64 `json:"employee_id" validate:"required"`
	LocationOut string `json:"location_out" validate:"required"`
}
