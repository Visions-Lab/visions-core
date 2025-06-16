/*
Copyright © 2025 Visions Lab
*/
package cronmgr

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/robfig/cron/v3"
)

// CronTask represents a scheduled job managed by CronManager.
//
// Fields:
//   - Name: unique name for the task
//   - Group: logical group for the task
//   - Spec: cron schedule string
//   - Command: command to execute
//   - Shell: whether to run the command in a shell
//
// ID is internal and not persisted.
type CronTask struct {
	ID      cron.EntryID `json:"-"`
	Name    string       `json:"name"`
	Group   string       `json:"group"`
	Spec    string       `json:"spec"`
	Command string       `json:"command"`
	Shell   bool         `json:"shell"`
}

// CronManager manages scheduled cron jobs with persistent storage.
type CronManager struct {
	cron     *cron.Cron          // The underlying cron scheduler
	tasks    map[string]CronTask // All tasks, keyed by name
	mu       sync.Mutex          // Mutex for thread safety
	filename string              // Path to persistent storage file
}

// NewCronManagerWithFile creates a new CronManager and loads tasks from the given file.
func NewCronManagerWithFile(filename string) *CronManager {
	m := &CronManager{
		cron:     cron.New(),
		tasks:    make(map[string]CronTask),
		filename: filename,
	}
	m.LoadTasks()
	return m
}

// Start begins the cron scheduler.
func (m *CronManager) Start() {
	m.cron.Start()
}

// AddTask adds or updates a cron task by name. If a task with the same name exists, it is replaced.
func (m *CronManager) AddTask(task CronTask) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Remove existing task if present
	if t, ok := m.tasks[task.Name]; ok {
		m.cron.Remove(t.ID)
	}
	id, err := m.cron.AddFunc(task.Spec, func() {
		cmd := buildCommand(task)
		if cmd != nil {
			cmd.Run()
		}
	})
	if err != nil {
		return err
	}
	task.ID = id
	m.tasks[task.Name] = task
	m.saveTasksLocked()
	return nil
}

// RemoveTask deletes a cron task by name.
func (m *CronManager) RemoveTask(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if t, ok := m.tasks[name]; ok {
		m.cron.Remove(t.ID)
		delete(m.tasks, name)
		m.saveTasksLocked()
	}
}

// RemoveGroup deletes all cron tasks in a given group.
func (m *CronManager) RemoveGroup(group string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	changed := false
	for name, t := range m.tasks {
		if t.Group == group {
			m.cron.Remove(t.ID)
			delete(m.tasks, name)
			changed = true
		}
	}
	if changed {
		m.saveTasksLocked()
	}
}

// ListTasks returns all cron tasks.
func (m *CronManager) ListTasks() []CronTask {
	return m.listTasksByGroup("")
}

// ListTasksByGroup returns all cron tasks in a given group.
func (m *CronManager) ListTasksByGroup(group string) []CronTask {
	return m.listTasksByGroup(group)
}

// listTasksByGroup is an internal helper to filter tasks by group.
func (m *CronManager) listTasksByGroup(group string) []CronTask {
	m.mu.Lock()
	defer m.mu.Unlock()
	result := []CronTask{}
	for _, t := range m.tasks {
		if group == "" || t.Group == group {
			result = append(result, t)
		}
	}
	return result
}

// saveTasksLocked writes all tasks to the persistent storage file. Caller must hold the lock.
func (m *CronManager) saveTasksLocked() {
	tasks := make([]CronTask, 0, len(m.tasks))
	for _, t := range m.tasks {
		tasks = append(tasks, t)
	}
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err == nil {
		_ = os.WriteFile(m.filename, data, 0644)
	}
}

// buildCommand creates an exec.Cmd based on the task's shell flag and command string.
func buildCommand(t CronTask) *exec.Cmd {
	if t.Shell {
		return exec.Command("sh", "-c", t.Command)
	}
	parts := strings.Fields(t.Command)
	if len(parts) > 0 {
		return exec.Command(parts[0], parts[1:]...)
	}
	return nil
}

// LoadTasks loads all tasks from the persistent storage file and registers them with the scheduler.
func (m *CronManager) LoadTasks() {
	m.mu.Lock()
	defer m.mu.Unlock()
	file, err := os.Open(m.filename)
	if err != nil {
		return // No file yet
	}
	defer file.Close()
	var tasks []CronTask
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return
	}
	for _, t := range tasks {
		id, err := m.cron.AddFunc(t.Spec, func() {
			cmd := buildCommand(t)
			if cmd != nil {
				cmd.Run()
			}
		})
		if err == nil {
			t.ID = id
			m.tasks[t.Name] = t
		}
	}
}
