# visions-core

[![Go Reference](https://pkg.go.dev/badge/github.com/Visions-Lab/visions-core/pkg/cronmgr.svg)](https://pkg.go.dev/github.com/Visions-Lab/visions-core/pkg/cronmgr)
[![CI](https://github.com/Visions-Lab/visions-core/actions/workflows/ci.yml/badge.svg)](https://github.com/Visions-Lab/visions-core/actions/workflows/ci.yml)

A professional, cross-platform Go CLI and API system for managing named and grouped cron jobs (scheduled tasks) in a modular, resource-efficient way.

## Features

- Add, update, list, and delete cron jobs by name or group
- Each job has a unique name, group, cron spec, and command (with optional shell)
- Persistent storage: jobs are saved to `cronjobs.json` and reloaded on startup
- Thread-safe, modular, and idiomatic Go code
- Well-documented and tested
- [See the Wiki for full documentation](https://github.com/Visions-Lab/visions-core/wiki)

## Quick Start

```sh
# Add a cron job
visions-core cron add --name=checkin --group=hoyoauto --spec="0 9 * * *" --exec="echo Hello" --shell

# List all jobs
visions-core cron list
```

## Go Module Usage

```go
import "github.com/Visions-Lab/visions-core/pkg/cronmgr"

mgr := cronmgr.NewCronManagerWithFile("cronjobs.json")
mgr.AddTask("hello", "demo", "* * * * *", "echo hello", true)
mgr.Start()
```

## Testing

```sh
go test ./...
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## License

MIT
