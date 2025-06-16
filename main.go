/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Visions-Lab/visions-core/cmd"
	"github.com/Visions-Lab/visions-core/pkg/cronmgr"
)

// main is the entry point for the Visions Core CLI application.
// It initializes the global cron manager, starts it, and runs the CLI.
func main() {
	cmd.Manager = cronmgr.NewCronManagerWithFile("cronjobs.json")
	cmd.Manager.Start()
	cmd.Execute()
}
