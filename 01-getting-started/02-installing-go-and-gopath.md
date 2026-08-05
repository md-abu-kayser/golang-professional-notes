# Installing Go & GOPATH

## Overview

This lesson covers installing Go on any major OS and understanding the environment variables Go relies on — including the historical `GOPATH` system and why modern Go (1.16+) mostly makes it invisible.

## Why It Matters

Environment setup problems are the single biggest source of early confusion — "it works in the tutorial but not on my machine" is almost always a `PATH` or module-mode issue.

## Core Concepts

### Installing Go

**macOS** (via Homebrew):
```bash
brew install go
```

**Linux** (manual install, adjust version as needed):
```bash
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

**Windows**: download and run the installer from [go.dev/dl](https://go.dev/dl) — it configures `PATH` automatically.

### Verifying the install

```bash
go version
# go version go1.22.0 linux/amd64
```

### GOROOT vs GOPATH vs modules

- **`GOROOT`** — where the Go installation itself lives (standard library, compiler). You rarely touch this.
- **`GOPATH`** — the historical (pre-2019) convention where *all* Go code had to live in one workspace (`$GOPATH/src/...`), and dependencies were fetched into `$GOPATH/pkg`. Every project shared one global set of dependency versions.
- **Go Modules** (default since Go 1.16) — each project declares its own dependencies and versions in a `go.mod` file, and can live *anywhere* on disk. This is the modern, correct approach — GOPATH-mode is effectively legacy.

```bash
go env GOPATH     # still exists — now just caches downloaded modules
go env GOROOT     # the Go installation directory
```

You'll interact with modules directly starting in [05-modules-and-packages-intro](./05-modules-and-packages-intro.md) — for now, just know: **new projects use `go mod init`, not `GOPATH` directories.**

## Common Pitfalls

- Following old tutorials that require code to live inside `$GOPATH/src` — unnecessary since Go 1.16; any directory works with modules.
- Installing Go via a package manager that lags several versions behind — check `go version` against [go.dev/dl](https://go.dev/dl) if something feels outdated.
- Forgetting to add Go's `bin` directory to `PATH` after a manual install, so `go` isn't recognized as a command.

## Key Takeaways

- `go version` confirms a working install.
- `GOROOT` is the Go installation itself; you rarely need to set it manually.
- `GOPATH` is a legacy workspace concept — modern Go uses per-project `go.mod` files instead.
- Any directory can be a Go project now; you are not confined to a single global workspace.
