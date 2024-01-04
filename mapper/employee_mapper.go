package mapper

import (
	"fmt"
	"github.com/okyirmawan/employee-attendance-api/domain"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func EmployeeRequestToDomain(dto dto.EmployeeRequest) domain.Employee {
	parsedDob, _ := time.Parse("2006-01-02", dto.DateOfBirth)
	gender, _ := ToGenderDomain(dto.Gender)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	return domain.Employee{
		Nip:         dto.Nip,
		Name:        dto.Name,
		DateOfBirth: parsedDob,
		Gender:      gender,
		Email:       dto.Email,
		Username:    dto.Username,
		Password:    string(hashedPassword),
		CreatedAt:   utils.CurrentLocalTime(),
		UpdatedAt:   utils.CurrentLocalTime(),
	}
}

func ToGenderDomain(genderStr string) (domain.Gender, error) {
	switch genderStr {
	case "Laki-laki":
		return domain.LakiLaki, nil
	case "Perempuan":
		return domain.Perempuan, nil
	default:
		return "", fmt.Errorf("unsupported gender: %s", genderStr)
	}
}

func ToEmployeeDomainList(dtos []dto.EmployeeRequest) []domain.Employee {
	Employees := make([]domain.Employee, len(dtos))

	for i, itm := range dtos {
		Employees[i] = EmployeeRequestToDomain(itm)
	}
	return Employees
}

func DomainToEmployeeResponse(Employee domain.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		Id:          Employee.ID,
		Nip:         Employee.Nip,
		Name:        Employee.Name,
		DateOfBirth: Employee.DateOfBirth.Format("2006-01-02"),
		Username:    Employee.Username,
		Email:       Employee.Email,
		Gender:      string(Employee.Gender),
	}
}

func DomainToEmployeeResponseList(employees []domain.Employee) []dto.EmployeeResponse {
	employeesDto := make([]dto.EmployeeResponse, len(employees))

	for i, itm := range employees {
		employeesDto[i] = DomainToEmployeeResponse(itm)
	}

	return employeesDto
}
