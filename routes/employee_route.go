package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/handler"
	"github.com/okyirmawan/employee-attendance-api/middlewares"
)

func EmployeeRoute(routes *echo.Echo, api handler.EmployeeAPI) {

	employeeMiddlewareRoute := routes.Group("/")
	employeeMiddlewareRoute.Use(middlewares.JwtAuthMiddleware)
	employeeMiddlewareRoute.PUT("employees", api.Update)
	employeeMiddlewareRoute.DELETE("employees", api.Delete)

	routes.POST("/employees/login", api.Login)
	routes.POST("/employees", api.Create)
	routes.GET("/employees", api.FindAll)
	routes.GET("/employees/:nip", api.FindByNip)
}
