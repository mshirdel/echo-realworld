package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/mshirdel/echo-realworld/app"
	api "github.com/mshirdel/echo-realworld/app/http"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	RunE:  serve,
}

func serve(_ *cobra.Command, _ []string) error {
	app := app.New(_cfgFile)
	err := app.InitAll()
	if err != nil {
		return err
	}

	defer app.Shutdown()

	server := api.NewHTTPServer(app)
	defer server.Shutdown()

	go server.Start()
	ctx, stop := handleInterrupts()
	defer stop()

	<-ctx.Done()
	return nil
}

func handleInterrupts() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
}
