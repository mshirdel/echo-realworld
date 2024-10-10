package cmd

import (
	"github.com/spf13/cobra"
)

var _cfgFile string

// _rootCmd represents the base command when called without any subcommands
var _rootCmd = &cobra.Command{
	Use:   "echo-realworld",
	Short: "A brief description of your application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the _rootCmd.
func Execute() error {
	return _rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	_rootCmd.PersistentFlags().StringVarP(&_cfgFile, "config", "c", "config.yaml", "config file")

	_rootCmd.AddCommand(serveCmd)
	_rootCmd.AddCommand(versionCmd)
	_rootCmd.AddCommand(_migrateCmd)
}
