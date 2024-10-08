package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/echo-realworld/app/service"
)

func AuthenticateUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth , ok := c.Request().Header["Authorization"]
			if !ok {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			parts := strings.Split(auth[0], " ")
			if len(parts) < 2 {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			if parts[0] != "Token" {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			if err := service.VerifyToken(parts[1]); err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			return next(c)
		}
	}
}
