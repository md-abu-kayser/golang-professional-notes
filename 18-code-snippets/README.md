# 18 — Code Snippets

> ✅ **Status:** Complete — 6/6 snippets

Standalone, `gofmt`-clean, copy-paste-ready `.go` files. Each is a runnable `package main` demonstrating one production pattern with inline comments explaining the *why*, not just the *what*.

| File | Pattern | Key APIs |
|---|---|---|
| [`error-handling-pattern.go`](./error-handling-pattern.go) | Sentinel errors, custom error types, wrapping | `errors.Is`, `errors.As`, `fmt.Errorf("%w", …)` |
| [`middleware-chain.go`](./middleware-chain.go) | Composable HTTP middleware | `http.Handler`, closures |
| [`worker-pool.go`](./worker-pool.go) | Bounded concurrency | channels, `sync.WaitGroup` |
| [`jwt-auth-middleware.go`](./jwt-auth-middleware.go) | Token-based auth middleware | `github.com/golang-jwt/jwt/v5`, `context.WithValue` |
| [`graceful-shutdown.go`](./graceful-shutdown.go) | Clean process shutdown | `os/signal`, `http.Server.Shutdown` |
| [`retry-with-backoff.go`](./retry-with-backoff.go) | Resilient external calls | exponential backoff + jitter |

## Running a snippet

```bash
go run error-handling-pattern.go
```

For `jwt-auth-middleware.go`, install the one external dependency first:

```bash
go mod init snippets && go get github.com/golang-jwt/jwt/v5
go run jwt-auth-middleware.go
```

All other snippets use only the standard library — no dependencies required.
