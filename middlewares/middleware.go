package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/okyirmawan/employee-attendance-api/utils/token"
	"net/http"
)

func JwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := token.TokenValid(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}
