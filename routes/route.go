package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db2 "github.com/okyirmawan/employee-attendance-api/db"
	"github.com/okyirmawan/employee-attendance-api/injection"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	dbConfig := db2.InitDB()

	employeeAPI := injection.InitEmployeeAPI(dbConfig)

	routes := echo.New()

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	EmployeeRoute(routes, employeeAPI)

	routes.GET("/swagger/*", echoSwagger.WrapHandler)

	return routes
}
