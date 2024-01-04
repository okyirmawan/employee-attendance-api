package dto

type AttendanceRequest struct {
	EmployeeId uint64 `json:"employee_id" validate:"required"`
	Latitude   string `json:"latitude" validate:"required"`
	Longitude  string `json:"longitude" validate:"required"`
}
