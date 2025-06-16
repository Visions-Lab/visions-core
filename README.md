# visions-core

[![Go Reference](https://pkg.go.dev/badge/github.com/Visions-Lab/visions-core/pkg/cronmgr.svg)](https://pkg.go.dev/github.com/Visions-Lab/visions-core/pkg/cronmgr)
[![CI](https://github.com/Visions-Lab/visions-core/actions/workflows/ci.yml/badge.svg)](https://github.com/Visions-Lab/visions-core/actions/workflows/ci.yml)

## What is visions-core?

**visions-core** is the minimal, production-grade foundation for the Visions ecosystem. It provides:

- A modular, thread-safe, and extensible Go core for automation, scheduling, and integration.
- A robust cron manager with persistent storage and CLI.
- A professional module registration system for building a scalable ecosystem.

**This repository is only the core.** All advanced features, modules, and documentation are maintained in separate repositories and the [visions-core wiki](https://github.com/Visions-Lab/visions-core/wiki).

## Design Philosophy

- Minimal, clean, and idiomatic Go code.
- Extensible by design: other repos can register modules and extend functionality.
- No bloat: only the essentials for a reliable, testable, and maintainable core.

## Quick Example

```sh
# Add a cron job
visions-core cron add --name=checkin --group=default --spec="0 9 * * *" --exec="echo Hello" --shell

# List all jobs
visions-core cron list
```

## Go Usage Example

```go
import "github.com/Visions-Lab/visions-core/pkg/cronmgr"

mgr := cronmgr.NewCronManagerWithFile("cronjobs.json")
mgr.AddTask(cronmgr.CronTask{
    Name:    "hello",
    Group:   "demo",
    Spec:    "* * * * *",
    Command: "echo hello",
    Shell:   true,
})
mgr.Start()
```

## Modules & Extensibility

visions-core supports dynamic module registration. Other repositories can implement and register modules to extend the core at runtime.

## License

MIT, check [LICENSE](LICENSE) for details.
