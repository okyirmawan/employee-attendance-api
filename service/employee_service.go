package service

import (
	"errors"
	"github.com/okyirmawan/employee-attendance-api/dto"
	"github.com/okyirmawan/employee-attendance-api/mapper"
	"github.com/okyirmawan/employee-attendance-api/repository"
	"github.com/okyirmawan/employee-attendance-api/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeRepositoryContract interface {
	Create(dto dto.EmployeeRequest) (dto.EmployeeResponse, error)
	FindAll() []dto.EmployeeResponse
	FindByNip(nim string) dto.EmployeeResponse
	Update(dto dto.EmployeeRequest, id uint64) (dto.EmployeeResponse, error)
	Delete(id string) error
	Auth(username string, password string) (dto.LoginResponse, error)
}

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepository
}

func ProviderEmployeeService(m repository.EmployeeRepository) EmployeeService {
	return EmployeeService{
		EmployeeRepository: m,
	}
}

func (m *EmployeeService) Create(dto dto.EmployeeRequest) (dto.EmployeeResponse, error) {
	employeeEntity := mapper.EmployeeRequestToDomain(dto)
	employee, err := m.EmployeeRepository.Create(employeeEntity)

	return mapper.DomainToEmployeeResponse(employee), err
}

func (m *EmployeeService) FindAll() []dto.EmployeeResponse {
	datas := m.EmployeeRepository.FindAll()

	return mapper.DomainToEmployeeResponseList(datas)
}

func (m *EmployeeService) FindByNip(nip string) dto.EmployeeResponse {
	employee := m.EmployeeRepository.FindByNip(nip)

	return mapper.DomainToEmployeeResponse(employee)
}

func (m *EmployeeService) Update(dto dto.EmployeeRequest, id uint64) (dto.EmployeeResponse, error) {
	employeeEntity := mapper.EmployeeRequestToDomain(dto)
	employee, err := m.EmployeeRepository.Update(employeeEntity, id)

	return mapper.DomainToEmployeeResponse(employee), err
}

func (m *EmployeeService) Delete(id uint64) error {
	employee := m.EmployeeRepository.FindByID(id)

	err := m.EmployeeRepository.Delete(employee)
	if err != nil {
		return err
	}

	return nil
}

func (m *EmployeeService) Auth(username string, password string) (dto.LoginResponse, error) {
	employee, err := m.EmployeeRepository.FindByUsername(username)
	loginResponse := dto.LoginResponse{}

	if err != nil {
		return loginResponse, err
	}

	err = VerifyPassword(password, employee.Password)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return loginResponse, err
	}

	authToken, err := token.GenerateToken(employee.ID)
	if err != nil {
		return loginResponse, err
	}

	loginResponse.Token = authToken
	return loginResponse, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
