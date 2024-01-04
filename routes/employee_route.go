package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/handler"
)

func EmployeeRoute(routes *echo.Echo, api handler.EmployeeAPI) {

	routes.POST("/employees", api.Create)
	routes.GET("/employees", api.FindAll)
	routes.GET("/employees/:nip", api.FindByNip)
	routes.PUT("/employees/:id", api.Update)
	routes.DELETE("/employees/:id", api.Delete)
}
