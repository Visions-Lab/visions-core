// Example Go usage for pkg/cronmgr
package main

import (
	"github.com/Visions-Lab/visions-core/pkg/cronmgr"
)

func main() {
	mgr := cronmgr.NewCronManagerWithFile("cronjobs.json")
	mgr.AddTask("hello", "demo", "* * * * *", "echo hello", true)
	mgr.Start()
}
