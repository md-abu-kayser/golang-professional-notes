# 02 — Language Basics

> 🚧 **Status:** Syllabus ready — 9 lessons outlined, full content in progress.

The core vocabulary of Go: how variables get their types, how the built-in control-flow constructs work, how functions are declared, and Go's distinctive approach to errors and cleanup. Everything downstream depends on this module being solid.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Variables Zero Values and Type Inference](./01-variables-zero-values-and-type-inference.md) | `var` vs `:=` declarations, the zero value every type gets automatically, and how type inference decides a variable's type. |
| 02 | [Basic Data Types](./02-basic-data-types.md) | Numeric types (`int`, `float64`, etc.), `bool`, `string`, their sizes, and explicit conversion rules between them. |
| 03 | [Constants and Iota](./03-constants-and-iota.md) | `const` blocks, the difference between typed and untyped constants, and the `iota` pattern for enumerations. |
| 04 | [Operators and Expressions](./04-operators-and-expressions.md) | Arithmetic, comparison, logical, and bitwise operators, plus operator precedence. |
| 05 | [Control Flow If Else Switch for](./05-control-flow-if-else-switch-for.md) | `if`/`else`, `switch` (including expressionless and fallthrough forms), and Go's single unified `for` loop. |
| 06 | [Functions Basics](./06-functions-basics.md) | Function declarations, multiple return values, variadic parameters, and named return values. |
| 07 | [Errors and Error Handling Pattern](./07-errors-and-error-handling-pattern.md) | The built-in `error` interface and Go's idiom of returning errors as ordinary values instead of throwing exceptions. |
| 08 | [Defer Panic Recover](./08-defer-panic-recover.md) | `defer`'s LIFO execution order for cleanup, and `panic`/`recover` for truly exceptional situations. |
| 09 | [Scope and Shadowing](./09-scope-and-shadowing.md) | Block-level scoping rules and the classic shadowing bug that `:=` can introduce inside nested blocks. |

**Previous module:** [01 — Getting Started](../01-getting-started/README.md) · **Next module:** [03 — Intermediate Types](../03-intermediate-types/README.md)
