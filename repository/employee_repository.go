package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/okyirmawan/employee-attendance-api/domain"
)

type EmployeeRepositoryContract interface {
	Create(employee domain.Employee) (domain.Employee, error)
	FindAll() []domain.Employee
	FindByID(id uint64) domain.Employee
	FindByNip(nip string) domain.Employee
	Update(employee domain.Employee, id int) (domain.Employee, error)
	Delete(employee domain.Employee) error
}

type EmployeeRepository struct {
	DB *gorm.DB
}

func ProviderEmployeeRepository(DB *gorm.DB) EmployeeRepository {
	return EmployeeRepository{DB: DB}
}

func (m *EmployeeRepository) Create(employee domain.Employee) (domain.Employee, error) {
	if err := m.DB.Create(&employee).Error; err != nil {
		return employee, err
	}
	return employee, nil
}

func (m *EmployeeRepository) FindAll() []domain.Employee {
	var employees []domain.Employee

	m.DB.Find(&employees)

	return employees
}

func (m *EmployeeRepository) FindByID(id uint64) domain.Employee {
	var employee domain.Employee

	m.DB.Where("id =? ", id).Find(&employee)

	return employee
}

func (m *EmployeeRepository) FindByNip(nip string) domain.Employee {
	var employee domain.Employee
	m.DB.Where("nip =? ", nip).Find(&employee)

	return employee
}

func (m *EmployeeRepository) Update(employee domain.Employee, id int) (domain.Employee, error) {
	employee.ID = uint64(id)

	if err := m.DB.Model(&domain.Employee{}).Where("id = ?", id).Updates(employee).Error; err != nil {
		return employee, err
	}

	return employee, nil
}

func (m *EmployeeRepository) Delete(employee domain.Employee) error {
	if err := m.DB.Delete(&employee).Error; err != nil {
		return err
	}
	return nil
}
