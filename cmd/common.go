/*
Copyright © 2025 Visions Lab
*/
package cmd

import "github.com/Visions-Lab/visions-core/pkg/cronmgr"

// Manager is the global cron manager instance shared by all commands.
// It must be initialized in main.go before any commands are run.
var Manager *cronmgr.CronManager
