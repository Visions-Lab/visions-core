package cronmgr

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type CronTask struct {
	ID    cron.EntryID
	Name  string
	Group string
	Spec  string
}

type CronManager struct {
	cron  *cron.Cron
	tasks map[string]CronTask // key: name
	mu    sync.Mutex
}

func NewCronManager() *CronManager {
	return &CronManager{
		cron:  cron.New(),
		tasks: make(map[string]CronTask),
	}
}

func (m *CronManager) Start() {
	m.cron.Start()
}

func (m *CronManager) AddTask(name, group, spec string, cmd func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	id, err := m.cron.AddFunc(spec, cmd)
	if err != nil {
		return err
	}
	m.tasks[name] = CronTask{ID: id, Name: name, Group: group, Spec: spec}
	return nil
}

func (m *CronManager) RemoveTask(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if t, ok := m.tasks[name]; ok {
		m.cron.Remove(t.ID)
		delete(m.tasks, name)
	}
}

func (m *CronManager) RemoveGroup(group string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for name, t := range m.tasks {
		if t.Group == group {
			m.cron.Remove(t.ID)
			delete(m.tasks, name)
		}
	}
}

func (m *CronManager) ListTasks() []CronTask {
	return m.listTasksByGroup("")
}

func (m *CronManager) ListTasksByGroup(group string) []CronTask {
	return m.listTasksByGroup(group)
}

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
