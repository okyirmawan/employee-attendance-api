package injection

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/okyirmawan/employee-attendance-api/handler"
	"github.com/okyirmawan/employee-attendance-api/repository"
	"github.com/okyirmawan/employee-attendance-api/service"
)

func initEmployeeAPI(db *gorm.DB) handler.EmployeeAPI {
	wire.Build(repository.ProviderEmployeeRepository, service.ProviderEmployeeService, handler.ProviderEmployeeAPI)

	return handler.EmployeeAPI{}
}
