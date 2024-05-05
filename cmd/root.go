package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "echo-realworld",
	Short: "A brief description of your application",
	// PersistentPreRunE: initConfig,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file")

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}
