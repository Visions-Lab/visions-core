// Package core provides the extensibility foundation for visions-core.
// It allows dynamic registration and management of modules.
package core

import (
	"sync"
)

// Module is the interface that all core modules must implement.
type Module interface {
	Name() string
	Init() error
}

var (
	moduleRegistry = make(map[string]Module)
	moduleMu       sync.RWMutex
)

// RegisterModule registers a module with the core system.
func RegisterModule(m Module) {
	moduleMu.Lock()
	defer moduleMu.Unlock()
	moduleRegistry[m.Name()] = m
}

// ListModules returns all registered modules.
func ListModules() []Module {
	moduleMu.RLock()
	defer moduleMu.RUnlock()
	mods := make([]Module, 0, len(moduleRegistry))
	for _, m := range moduleRegistry {
		mods = append(mods, m)
	}
	return mods
}

// GetModule returns a module by name, or nil if not found.
func GetModule(name string) Module {
	moduleMu.RLock()
	defer moduleMu.RUnlock()
	return moduleRegistry[name]
}

// BuiltinModule is an example module for demonstration.
type BuiltinModule struct{}

func (m *BuiltinModule) Name() string { return "builtin" }
func (m *BuiltinModule) Init() error  { return nil }
