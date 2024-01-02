package dto

type EmployeeDTO struct {
	Id          uint64 `json:"id"`
	Nip         string `json:"nip" validate:"required"`
	Name        string `json:"name" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
}
