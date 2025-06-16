package cmd

import "github.com/Visions-Lab/visions-core/internal/cronmgr"

// Manager is the global cron manager instance shared by all commands.
var Manager *cronmgr.CronManager
