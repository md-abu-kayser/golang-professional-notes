# 14 — Performance Optimization

> 🚧 **Status:** Syllabus ready — 6 lessons outlined, full content in progress.

Measuring before optimizing: profiling tools, memory and GC behavior, and the handful of patterns that actually move the needle in Go programs.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Profiling with Pprof](./01-profiling-with-pprof.md) | Finding real bottlenecks with `pprof` before touching a single line of code. |
| 02 | [Memory Allocation and Escape Analysis](./02-memory-allocation-and-escape-analysis.md) | Reading `go build -gcflags='-m'` escape analysis output to understand allocation. |
| 03 | [Garbage Collection Tuning](./03-garbage-collection-tuning.md) | How Go's garbage collector behaves, and what `GOGC` actually controls. |
| 04 | [Benchmarking and Comparing Algorithms](./04-benchmarking-and-comparing-algorithms.md) | Writing statistically meaningful benchmarks with `go test -bench`. |
| 05 | [String Concatenation and Immutable Optimizations](./05-string-concatenation-and-immutable-optimizations.md) | Why `strings.Builder` beats repeated `+=` concatenation in loops. |
| 06 | [Concurrency Optimization Patterns](./06-concurrency-optimization-patterns.md) | Avoiding goroutine and channel overhead pitfalls at scale. |

**Previous module:** [13 — Security](../13-security/README.md) · **Next module:** [15 — Deployment](../15-deployment/README.md)
