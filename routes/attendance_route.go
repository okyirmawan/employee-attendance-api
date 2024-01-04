package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/handler"
	"github.com/okyirmawan/employee-attendance-api/middlewares"
)

func AttendanceRoute(routes *echo.Echo, api handler.AttendanceAPI) {

	attendanceMiddlewareRoute := routes.Group("/attendances")
	attendanceMiddlewareRoute.Use(middlewares.JwtAuthMiddleware)

	attendanceMiddlewareRoute.POST("/checkin", api.CheckIn)
	attendanceMiddlewareRoute.PUT("/checkout/:id", api.CheckOut)
	attendanceMiddlewareRoute.GET("/history", api.History)
}
