package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/echo-realworld/app"
)

type Controller struct {
	app  *app.Application
	user *UserController
}

func NewController(app *app.Application) *Controller {
	return &Controller{
		app:  app,
		user: NewUserController(app.Svc),
	}
}

func (c *Controller) Routes() *echo.Echo {
	router := c.initEcho()
	router.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, c.app.Cfg.Database.Logger.Level)
	})

	api := router.Group("/api")
	{
		api.GET("/users", c.user.GetUsers)
		api.POST("/users", c.user.RegisterUser)
	}

	return router
}

func (c *Controller) initEcho() *echo.Echo {
	e := echo.New()
	e.Debug = c.app.Cfg.Logging.Level == "debug"

	return e
}
