# Modules & Packages Intro

## Overview

Go organizes code into **packages** (units of compiled code) grouped into **modules** (versioned collections of packages with declared dependencies). This lesson covers the two files — `go.mod` and `go.sum` — that define a module, and the naming conventions around packages.

## Why It Matters

Every real Go project — including every project in [17-real-world-projects](../17-real-world-projects/README.md) — starts with `go mod init`. Understanding what that command actually creates removes a lot of "magic" from the rest of the curriculum.

## Core Concepts

### Creating a module

```bash
go mod init github.com/md-abu-kayser/my-project
```

This creates `go.mod`:

```text
module github.com/md-abu-kayser/my-project

go 1.22
```

- **`module`** — the import path other code will use to reference this module.
- **`go 1.22`** — the minimum Go version this module requires.

### Adding dependencies

```bash
go get github.com/gin-gonic/gin
```

This updates `go.mod` with a `require` line, and generates/updates `go.sum` — a lockfile of exact, cryptographically verified dependency versions:

```text
require github.com/gin-gonic/gin v1.9.1
```

### Keeping dependencies clean

```bash
go mod tidy
```

Adds anything imported-but-missing from `go.mod`, and removes anything listed-but-unused. Run this before every commit that touches imports.

### Packages and naming

- Every `.go` file starts with a `package` declaration; all files in the same directory must share the same package name.
- Package names are short, lowercase, no underscores: `http`, `json`, `strings` — not `HttpUtils` or `json_helpers`.
- The directory name and package name are conventionally the same, though not required to be.
- `package main` is reserved for executables; everything else is an importable library package.

```go
// file: userservice/user.go
package userservice

func New() *Service { /* ... */ }
```

```go
// file: main.go
package main

import "github.com/md-abu-kayser/my-project/userservice"

func main() {
    svc := userservice.New()
    _ = svc
}
```

## Common Pitfalls

- Editing `go.mod` by hand instead of using `go get`/`go mod tidy` — easy to end up with an inconsistent `go.sum`.
- Choosing a module path that doesn't match where the code actually lives on GitHub — breaks `go get` for anyone trying to import it.
- Package names that stutter with their contents, e.g. `httputil.HTTPClient` — idiomatic Go drops the redundant prefix: `httputil.Client`.

## Key Takeaways

- `go.mod` declares the module path, Go version, and dependencies; `go.sum` locks exact dependency checksums.
- `go get` adds/updates a dependency; `go mod tidy` reconciles imports with what's declared.
- Package names are short and lowercase; `package main` is reserved for executables.
- The module path should match the repository URL so others can `go get` it directly.
