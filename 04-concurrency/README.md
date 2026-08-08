# 04 — Concurrency

> 🚧 **Status:** Syllabus ready — 10 lessons outlined, full content in progress.

Go's headline feature: concurrency as a language primitive rather than a library. This module covers goroutines, channels, and the patterns and pitfalls that come with building concurrent systems — see also the working [worker-pool.go](../18-code-snippets/worker-pool.go) snippet.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Goroutines](./01-goroutines.md) | Lightweight goroutines, the `go` keyword, and a plain-language overview of Go's M:N scheduler. |
| 02 | [Channels Unbuffered and Buffered](./02-channels-unbuffered-and-buffered.md) | Channel semantics, blocking behavior, and the difference between buffered and unbuffered channels. |
| 03 | [Select Statement](./03-select-statement.md) | Multiplexing across multiple channel operations with `select`, including `default` and timeout patterns. |
| 04 | [Mutex and RWMutex](./04-mutex-and-rwmutex.md) | Protecting shared state with `sync.Mutex` and `sync.RWMutex` when channels aren't the right tool. |
| 05 | [WaitGroup and ErrGroup](./05-waitgroup-and-errgroup.md) | Coordinating goroutine completion with `sync.WaitGroup` and error propagation with `golang.org/x/sync/errgroup`. |
| 06 | [Concurrency Patterns Pipelines Fan in Fan Out](./06-concurrency-patterns-pipelines-fan-in-fan-out.md) | Composable concurrency patterns: pipeline stages, fan-out, and fan-in. |
| 07 | [Context Package Cancellation Deadlines](./07-context-package-cancellation-deadlines.md) | `context.Context` for cancellation signals, deadlines, and request-scoped values — see also [graceful-shutdown.go](../18-code-snippets/graceful-shutdown.go). |
| 08 | [Atomic Operations](./08-atomic-operations.md) | Lock-free counters and flags with the `sync/atomic` package. |
| 09 | [Race Detector](./09-race-detector.md) | Catching data races before production with `go run -race` and `go test -race`. |
| 10 | [Common Concurrency Pitfalls](./10-common-concurrency-pitfalls.md) | Goroutine leaks, the classic loop-variable-capture bug, and deadlock patterns to recognize. |

**Previous module:** [03 — Intermediate Types](../03-intermediate-types/README.md) · **Next module:** [05 — Standard Library](../05-standard-library/README.md)
