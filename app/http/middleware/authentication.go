package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/echo-realworld/app/service"
)

const _authHeader = "Authorization"

func AuthenticateUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get(_authHeader)
			if len(auth) == 0 {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			parts := strings.Split(auth, " ")
			if len(parts) < 2 {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			if parts[0] != "Token" {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			userID, err := service.VerifyToken(parts[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			if userID != nil {
				c.Set("user_id", userID)
			}

			return next(c)
		}
	}
}
