package cronmgr_test

import (
	"os"
	"testing"

	"github.com/Visions-Lab/visions-core/pkg/cronmgr"
)

func TestAddListRemoveTask(t *testing.T) {
	os.Remove("test_cronjobs.json")
	mgr := cronmgr.NewCronManagerWithFile("test_cronjobs.json")
	defer os.Remove("test_cronjobs.json")

	err := mgr.AddTask("t1", "g1", "* * * * *", "echo hi", true)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}
	tasks := mgr.ListTasks()
	if len(tasks) != 1 || tasks[0].Name != "t1" {
		t.Fatalf("Expected 1 task named t1, got %+v", tasks)
	}

	mgr.RemoveTask("t1")
	if len(mgr.ListTasks()) != 0 {
		t.Fatalf("Expected 0 tasks after remove")
	}
}

func TestPersistence(t *testing.T) {
	os.Remove("test_cronjobs.json")
	mgr := cronmgr.NewCronManagerWithFile("test_cronjobs.json")
	mgr.AddTask("t2", "g2", "* * * * *", "echo hi", true)
	mgr = nil

	mgr2 := cronmgr.NewCronManagerWithFile("test_cronjobs.json")
	tasks := mgr2.ListTasks()
	if len(tasks) != 1 || tasks[0].Name != "t2" {
		t.Fatalf("Persistence failed, got %+v", tasks)
	}
	mgr2.RemoveTask("t2")
	os.Remove("test_cronjobs.json")
}
