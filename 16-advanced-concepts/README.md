# 16 — Advanced Concepts

> 🚧 **Status:** Syllabus ready — 6 lessons outlined, full content in progress.

The corners of Go most engineers touch rarely but should understand: generics, reflection, cgo, and how the compiler and runtime actually work under the hood.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Generics and Type Parameters](./01-generics-and-type-parameters.md) | Type parameters and constraints, and a clear-eyed view of when generics genuinely help. |
| 02 | [Reflection and Unsafe Package](./02-reflection-and-unsafe-package.md) | Runtime type inspection with `reflect`, and why `unsafe` is a last resort, not a shortcut. |
| 03 | [Cgo and Interfacing with C](./03-cgo-and-interfacing-with-c.md) | Calling C code from Go with `cgo`, and its build/performance tradeoffs. |
| 04 | [Plugin Architecture and Rpc](./04-plugin-architecture-and-rpc.md) | Go's `plugin` package and `net/rpc` for pluggable architectures. |
| 05 | [Design Patterns in Go](./05-design-patterns-in-go.md) | Idiomatic Go takes on classic design patterns — where they translate, and where Go prefers a simpler idiom instead. |
| 06 | [Go Compiler and Runtime Internals](./06-go-compiler-and-runtime-internals.md) | A guided look at how compilation, the scheduler, and the runtime actually work. |

**Previous module:** [15 — Deployment](../15-deployment/README.md) · **Next module:** [17 — Real-World Projects](../17-real-world-projects/README.md)
