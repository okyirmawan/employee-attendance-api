package mapper

import (
	"fmt"
	"github.com/okyirmawan/employee-attendance-api/domain"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/utils"
	"time"
)

func ToEmployeeDomain(dto dto.EmployeeDTO) domain.Employee {
	parsedDob, _ := time.Parse("2006-01-02", dto.DateOfBirth)
	gender, _ := ToGenderDomain(dto.Gender)

	return domain.Employee{
		Nip:         dto.Nip,
		Name:        dto.Name,
		DateOfBirth: parsedDob,
		Gender:      gender,
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

func ToEmployeeDomainList(dtos []dto.EmployeeDTO) []domain.Employee {
	Employees := make([]domain.Employee, len(dtos))

	for i, itm := range dtos {
		Employees[i] = ToEmployeeDomain(itm)
	}
	return Employees
}

func ToEmployeeDto(Employee domain.Employee) dto.EmployeeDTO {
	return dto.EmployeeDTO{
		Id:          Employee.ID,
		Nip:         Employee.Nip,
		Name:        Employee.Name,
		DateOfBirth: Employee.DateOfBirth.Format("2006-01-02"),
		Gender:      string(Employee.Gender),
	}
}

func ToEmployeeDtoList(Employees []domain.Employee) []dto.EmployeeDTO {
	dtos := make([]dto.EmployeeDTO, len(Employees))

	for i, itm := range Employees {
		dtos[i] = ToEmployeeDto(itm)
	}

	return dtos
}
