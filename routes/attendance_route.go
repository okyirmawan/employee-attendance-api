package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/handler"
)

func AttendanceRoute(routes *echo.Echo, api handler.AttendanceAPI) {
	routes.POST("/attendance/checkin", api.CheckIn)
	routes.PUT("/attendance/checkout/:id", api.CheckOut)
	routes.GET("/attendance/history/:employee_id", api.History)
}
