package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/echo-realworld/app"
	"github.com/mshirdel/echo-realworld/model"
)

type Controller struct {
	app *app.Application
}

func NewController(app *app.Application) *Controller {
	return &Controller{
		app: app,
	}
}

func (c *Controller) Routes() *echo.Echo {
	router := c.initEcho()
	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	router.GET("/users", func(c echo.Context) error {
		u := model.User{
			Username: "mshirdel",
			Email: "mshirdel@gmail.com",
		}
		
		return c.JSON(http.StatusOK, u)
	})

	return router
}

func (c *Controller) initEcho() *echo.Echo {
	e := echo.New()
	e.Debug = c.app.Cfg.Logging.Level == "debug"

	return e
}
