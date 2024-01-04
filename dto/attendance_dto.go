package dto

type AttendanceDTO struct {
	Id          uint64 `json:"id"`
	EmployeeId  uint64 `json:"employee_id"`
	CheckIn     string `json:"check_in"`
	CheckOut    string `json:"check_out"`
	LocationIn  string `json:"location_in"`
	LocationOut string `json:"location_out"`
}
