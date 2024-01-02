// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/jinzhu/gorm"
	"github.com/okyirmawan/employee-attendance-api/handler"
	"github.com/okyirmawan/employee-attendance-api/repository"
	"github.com/okyirmawan/employee-attendance-api/service"
)

// Injectors from wire.go:
func InitEmployeeAPI(db *gorm.DB) handler.EmployeeAPI {
	employeeRepository := repository.ProviderEmployeeRepository(db)
	employeeService := service.ProviderEmployeeService(employeeRepository)
	employeeAPI := handler.ProviderEmployeeAPI(employeeService)
	return employeeAPI
}
