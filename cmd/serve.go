package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"github.com/mshirdel/echo-realworld/internal/config"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serve,
}

func serve(_ *cobra.Command, _ []string) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(config.C.HTTPServer.Address))
}
