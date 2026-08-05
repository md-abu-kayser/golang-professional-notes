# Go Commands: run, build, install, fmt

## Overview

A cheat-sheet-style lesson on the Go CLI commands you'll type dozens of times a day. Go's tooling is deliberately unified — one `go` binary handles building, testing, formatting, dependency management, and documentation.

## Why It Matters

Fluency with these commands is what separates "I can write Go" from "I can ship Go" — CI pipelines, Makefiles, and Dockerfiles across the industry are built from exactly these commands.

## Core Concepts

### `go run` — compile and execute, no artifact kept

```bash
go run main.go
go run .              # run the main package in the current directory
```
Best for quick iteration; discards the compiled binary afterward.

### `go build` — compile to a binary, don't run it

```bash
go build .                    # produces a binary named after the module/dir
go build -o myapp .             # explicit output name
GOOS=linux GOARCH=amd64 go build .  # cross-compile (see 15-deployment)
```

### `go install` — build and place the binary on your `PATH`

```bash
go install .
# binary lands in $(go env GOPATH)/bin — runnable from anywhere if that's on PATH
```
Commonly used to install your own CLI tools, or third-party ones (`go install github.com/user/tool@latest`).

### `go fmt` — canonical formatting, non-negotiable

```bash
go fmt ./...     # reformat every file in the module
gofmt -l .        # list files that are NOT correctly formatted (used in CI)
```
Go has one canonical style. This ends formatting debates in code review permanently.

### Other daily commands

```bash
go vet ./...       # static analysis — catches likely bugs (see 07-tooling)
go test ./...        # run tests (see 06-testing)
go doc fmt.Println     # read documentation without leaving the terminal
```

## Common Pitfalls

- Committing unformatted code — most teams enforce `gofmt` in CI, so this fails builds, not just style checks.
- Confusing `go build` (produces a binary) with `go install` (produces a binary *and* places it on `PATH`) — using the wrong one for CLI tools leads to "why can't I run this from anywhere" confusion.
- Forgetting `./...` — running `go vet` or `go test` without it only checks the current directory, not sub-packages.

## Key Takeaways

- `go run` = compile + execute + discard; `go build` = compile + keep binary; `go install` = compile + place on `PATH`.
- `go fmt` / `gofmt` is the single source of truth for code style — run it before every commit.
- `./...` means "this package and every sub-package" — use it for project-wide commands.
- `go vet` and `go test` are part of the same unified toolchain, not separate installs.
