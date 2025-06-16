/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Visions-Lab/visions-core/cmd"
	"github.com/Visions-Lab/visions-core/internal/cronmgr"
)

func main() {
	cmd.Manager = cronmgr.NewCronManager()
	cmd.Manager.Start()
	cmd.Execute()
}
