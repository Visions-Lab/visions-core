/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Visions-Lab/visions-core/pkg/core"
	"github.com/sirupsen/logrus"

	"os"

	"github.com/Visions-Lab/visions-core/cmd"
	"github.com/Visions-Lab/visions-core/pkg/config"
	"github.com/Visions-Lab/visions-core/pkg/cronmgr"
)

// main is the entry point for the Visions Core CLI application.
// It initializes the global cron manager, starts it, and runs the CLI.
func main() {
	// Example: Register a built-in module (other repos can do this too)
	core.RegisterModule(&core.BuiltinModule{})

	cfg := mustLoadConfig("config.json")
	setupLogging(cfg)
	cronFile := getCronFile(cfg)
	cmd.Manager = cronmgr.NewCronManagerWithFile(cronFile)
	cmd.Manager.Start()
	cmd.Execute()
}

func mustLoadConfig(path string) *config.AppConfig {
	cfg, err := config.Load(path)
	if err != nil {
		logrus.Warnf("Could not load %s, using default cron file. Error: %v", path, err)
	}
	return cfg
}

func getCronFile(cfg *config.AppConfig) string {
	if cfg != nil && cfg.CronFile != "" {
		return cfg.CronFile
	}
	return "cronjobs.json"
}

func setupLogging(cfg *config.AppConfig) {
	if cfg == nil {
		logrus.SetLevel(logrus.InfoLevel)
		return
	}
	level := logrus.InfoLevel
	if cfg.LogLevel != "" {
		if parsed, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
			level = parsed
		}
	}
	logrus.SetLevel(level)
	if cfg.LogFile == "" {
		return
	}
	f, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		logrus.SetOutput(f)
	} else {
		logrus.Warnf("Could not open log file %s: %v", cfg.LogFile, err)
	}
}
