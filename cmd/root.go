/*
Copyright © 2025 Visions Lab
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "visions-core",
	Short: "Modular core framework for automation and integration",
	Long: `Visions Core is the modular, resource-efficient foundation 
for automation, integration, and workflow tools in the Visions Lab ecosystem.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

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

func init() {
	// Define persistent flags (global for all commands) here if needed:
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.visions-core.yaml)")

	// Remove example toggle flag unless you plan to use it:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(startCmd)
}
