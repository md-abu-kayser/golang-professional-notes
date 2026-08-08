# 07 — Tooling & Code Quality

> 🚧 **Status:** Syllabus ready — 7 lessons outlined, full content in progress.

The tools that keep a Go codebase consistent and catch bugs before code review: formatters, static analyzers, linters, debuggers, and profilers.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Go Fmt and Goimports](./01-go-fmt-and-goimports.md) | Canonical formatting with `gofmt`, and auto-organizing imports with `goimports`. |
| 02 | [Go Vet and Staticcheck](./02-go-vet-and-staticcheck.md) | Catching likely bugs before runtime with `go vet` and the more thorough `staticcheck`. |
| 03 | [Golangci Lint Setup](./03-golangci-lint-setup.md) | Configuring `golangci-lint` to run dozens of linters together in CI. |
| 04 | [Debugging with Delve](./04-debugging-with-delve.md) | Step-through debugging, breakpoints, and inspecting goroutines with the `dlv` debugger. |
| 05 | [Pprof and Performance Profiling](./05-pprof-and-performance-profiling.md) | CPU and memory profiling with `net/http/pprof` and the `go tool pprof` viewer. |
| 06 | [Race Detector in Depth](./06-race-detector-in-depth.md) | How the race detector actually works under the hood, and its coverage limitations. |
| 07 | [Cross Compilation and Build Tags](./07-cross-compilation-and-build-tags.md) | Cross-compiling with `GOOS`/`GOARCH`, and conditional compilation with build tags. |

**Previous module:** [06 — Testing](../06-testing/README.md) · **Next module:** [08 — Modules & Dependency Management](../08-modules-and-dependency-management/README.md)
