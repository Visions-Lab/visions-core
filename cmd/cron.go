/*
Copyright © 2025 Visions Lab
*/
package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/Visions-Lab/visions-core/pkg/cronmgr"

	"github.com/spf13/cobra"
)

// Flags for cron task operations
var (
	// taskName is the unique name for the cron task
	taskName string
	// taskGroup is the logical group for the cron task
	taskGroup string
	// taskSpec is the cron schedule string (e.g. "* * * * *")
	taskSpec string
	// taskCommand is the command to execute for the cron task
	taskCommand string
	// taskShell determines if the command should be run in a shell
	taskShell bool
)

// cronCmd is the parent command for all cron-related subcommands
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Manage cron tasks",
	Long:  `Manage and schedule cron tasks for Visions Core.`,
}

// cronAddCmd adds or updates a cron task
var cronAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new cron task",
	Example: `visions-core cron add --name mytask --group default --spec "*/5 * * * *" --exec "echo hello" --shell`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if taskName == "" || taskGroup == "" || taskSpec == "" || taskCommand == "" {
			logrus.Error("You must provide --name, --group, --spec, and --exec.")
			return fmt.Errorf("invalid arguments: missing required flags")
		}
		task := cronmgr.CronTask{
			Name:    taskName,
			Group:   taskGroup,
			Spec:    taskSpec,
			Command: taskCommand,
			Shell:   taskShell,
		}
		err := Manager.AddTask(task)
		if err != nil {
			logrus.Errorf("Failed to add/update task: %v", err)
			return err
		}
		logrus.Info("Task added or updated successfully!")
		return nil
	},
}

// cronListCmd lists all cron tasks, optionally filtered by group
var cronListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cron tasks",
	Long:  `List all cron tasks, or filter by group.`,
	Run: func(cmd *cobra.Command, args []string) {
		group, _ := cmd.Flags().GetString("group")
		var tasks []interface{}
		if group == "" {
			tasks = make([]interface{}, 0, len(Manager.ListTasks()))
			for _, t := range Manager.ListTasks() {
				tasks = append(tasks, t)
			}
		} else {
			tasks = make([]interface{}, 0, len(Manager.ListTasksByGroup(group)))
			for _, t := range Manager.ListTasksByGroup(group) {
				tasks = append(tasks, t)
			}
		}
		if len(tasks) == 0 {
			logrus.Info("No cron tasks found.")
			return
		}
		logrus.Info("Name\tGroup\tSpec\tShell\tCommand")
		for _, t := range tasks {
			task := t.(interface {
				Name() string
				Group() string
				Spec() string
				Shell() bool
				Command() string
			})
			logrus.Infof("%s\t%s\t%s\t%v\t%s", task.Name(), task.Group(), task.Spec(), task.Shell(), task.Command())
		}
	},
}

// cronDelCmd deletes a cron task by name or all tasks in a group
var cronDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete cron tasks by name or group",
	Long:  `Delete a cron task by name, or delete all tasks in a group.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		group, _ := cmd.Flags().GetString("group")
		if name == "" && group == "" {
			logrus.Error("Provide --name or --group to delete.")
			return fmt.Errorf("invalid arguments: must provide --name or --group")
		}
		if name != "" {
			Manager.RemoveTask(name)
			logrus.Infof("Task '%s' deleted (if it existed).", name)
		}
		if group != "" {
			Manager.RemoveGroup(group)
			logrus.Infof("All tasks in group '%s' deleted (if any existed).", group)
		}
		return nil
	},
}

// init registers all cron subcommands and flags
func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.AddCommand(cronAddCmd)
	cronAddCmd.Flags().StringVar(&taskName, "name", "", "Task name (required)")
	cronAddCmd.Flags().StringVar(&taskGroup, "group", "", "Task group (required)")
	cronAddCmd.Flags().StringVar(&taskSpec, "spec", "", "Cron spec (e.g. '*/5 * * * *') (required)")
	cronAddCmd.Flags().StringVar(&taskCommand, "exec", "", "Command to execute (required)")
	cronAddCmd.Flags().BoolVar(&taskShell, "shell", false, "Run command in shell (default: false)")
	cronAddCmd.MarkFlagRequired("name")
	cronAddCmd.MarkFlagRequired("group")
	cronAddCmd.MarkFlagRequired("spec")
	cronAddCmd.MarkFlagRequired("exec")

	cronCmd.AddCommand(cronListCmd)
	cronListCmd.Flags().String("group", "", "Filter by group")
	cronCmd.AddCommand(cronDelCmd)
	cronDelCmd.Flags().String("name", "", "Task name to delete")
	cronDelCmd.Flags().String("group", "", "Group to delete all tasks from")
}
