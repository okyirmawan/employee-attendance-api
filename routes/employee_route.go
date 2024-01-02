package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/handler"
)

func EmployeeRoute(routes *echo.Echo, api handler.EmployeeAPI) {

	routes.POST("/employee", api.Create)
	routes.GET("/employees", api.FindAll)
	routes.GET("/employee/nip/:nip", api.FindByNip)
	routes.PUT("/employee/:id", api.Update)
}
