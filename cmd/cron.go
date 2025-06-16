/*
Copyright © 2025 Visions Lab
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	taskName  string
	taskGroup string
	taskSpec  string
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Manage cron tasks",
	Long:  `Manage and schedule cron tasks for Visions Core.`,
}

var cronAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new cron task",
	Example: `visions-core cron add --name mytask --group default --spec "*/5 * * * *"`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskName == "" || taskGroup == "" || taskSpec == "" {
			fmt.Fprintln(os.Stderr, "Error: You must provide --name, --group, and --spec.")
			os.Exit(1)
		}
		// Use the exported Manager
		err := Manager.AddTask(taskName, taskGroup, taskSpec, func() {
			fmt.Printf("Task %s in group %s executed!\n", taskName, taskGroup)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add task: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.AddCommand(cronAddCmd)
	cronAddCmd.Flags().StringVar(&taskName, "name", "", "Task name (required)")
	cronAddCmd.Flags().StringVar(&taskGroup, "group", "", "Task group (required)")
	cronAddCmd.Flags().StringVar(&taskSpec, "spec", "", "Cron spec (e.g. '*/5 * * * *') (required)")
	cronAddCmd.MarkFlagRequired("name")
	cronAddCmd.MarkFlagRequired("group")
	cronAddCmd.MarkFlagRequired("spec")

	// Future: Add more subcommands (list, remove, update, etc.)
}
