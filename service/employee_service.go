package service

import (
	"errors"
	"github.com/okyirmawan/employee-attendance-api/domain"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/mapper"
	"github.com/okyirmawan/employee-attendance-api/repository"
)

type EmployeeRepositoryContract interface {
	Create(dto dto.EmployeeDTO) (dto.EmployeeDTO, error)
	FindAll() []dto.EmployeeDTO
	FindByNip(nim string) dto.EmployeeDTO
	Update(dto dto.EmployeeDTO, id int) (dto.EmployeeDTO, error)
	DeleteEmployee(id string) error
}

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepository
}

func ProviderEmployeeService(m repository.EmployeeRepository) EmployeeService {
	return EmployeeService{
		EmployeeRepository: m,
	}
}

func (m *EmployeeService) Create(dto dto.EmployeeDTO) (dto.EmployeeDTO, error) {
	employeeEntity := mapper.ToEmployeeDomain(dto)
	employee, err := m.EmployeeRepository.Create(employeeEntity)

	return mapper.ToEmployeeDto(employee), err
}

func (m *EmployeeService) FindAll() []dto.EmployeeDTO {
	datas := m.EmployeeRepository.FindAll()

	return mapper.ToEmployeeDtoList(datas)
}

func (m *EmployeeService) FindByNip(nip string) dto.EmployeeDTO {
	employee := m.EmployeeRepository.FindByNip(nip)

	return mapper.ToEmployeeDto(employee)
}

func (m *EmployeeService) Update(dto dto.EmployeeDTO, id int) (dto.EmployeeDTO, error) {
	employeeEntity := mapper.ToEmployeeDomain(dto)
	employee, err := m.EmployeeRepository.Update(employeeEntity, id)

	return mapper.ToEmployeeDto(employee), err
}

func (m *EmployeeService) DeleteEmployee(id string) error {

	employee := m.EmployeeRepository.FindByID(id)

	if employee == (domain.Employee{}) {
		return errors.New("Employee not found")
	}

	err := m.EmployeeRepository.DeleteEmployee(employee)
	if err != nil {
		return err
	}

	return nil
}
