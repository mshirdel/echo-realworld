package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mshirdel/echo-realworld/app"
	"github.com/mshirdel/echo-realworld/app/http/controller"
	"github.com/sirupsen/logrus"
)

type Server struct {
	app    *app.Application
	server *http.Server
}

func NewHTTPServer(app *app.Application) *Server {
	controller := controller.NewController(app)

	return &Server{
		app: app,
		server: &http.Server{
			Addr:         app.Cfg.HTTPServer.Address,
			ReadTimeout:  app.Cfg.HTTPServer.ReadTimeout,
			WriteTimeout: app.Cfg.HTTPServer.WriteTimeout,
			IdleTimeout:  app.Cfg.HTTPServer.IdleTimeout,
			Handler:      controller.Routes(),
		},
	}
}

func (s *Server) Start() {
	logrus.Infof("starting http server on: %s\n", s.app.Cfg.HTTPServer.Address)
	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("faild starting http server: %v", err)
	}
}

func (s *Server) Shutdown() {
	logrus.Infof("shutting down http server...")
	deadline, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(deadline); err != nil {
		logrus.Errorf("faild shutting down http server: %v", err)
	}
}
