/*
Copyright © 2025 Visions Lab
*/
package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd is the base command for the CLI, used when no subcommands are provided.
var rootCmd = &cobra.Command{
	Use:   "visions-core",
	Short: "Modular core framework for automation and integration",
	Long: `Visions Core is the modular, resource-efficient foundation 
for automation, integration, and workflow tools in the Visions Lab ecosystem.`,
}

// Execute runs the root command and all subcommands. Exits with code 1 on error.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

// startCmd starts all background services (cron, etc) and blocks until interrupted.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Visions Core services (cron, etc)",
	Long:  `Start all background services and schedulers for Visions Core.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Manager == nil {
			cmd.PrintErrln("Error: Manager is not initialized.")
			os.Exit(1)
		}
		Manager.Start()
		cmd.Println("Visions Core services started. Press Ctrl+C to exit.")
		select {} // Block forever
	},
}

// init registers the start command and any persistent flags with the root command.
func init() {
	// You can define persistent flags (global for all commands) here if needed:
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.visions-core.yaml)")

	// Remove example toggle flag unless you plan to use it:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(startCmd)
}
