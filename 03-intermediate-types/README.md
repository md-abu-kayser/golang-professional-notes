# 03 — Intermediate Types

> 🚧 **Status:** Syllabus ready — 9 lessons outlined, full content in progress.

Go's composite types — arrays, slices, maps, structs, and pointers — and the two mechanisms (interfaces and embedding) it uses instead of classical inheritance.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Arrays](./01-arrays.md) | Fixed-size arrays, their value (copy) semantics, and why slices are used far more often in practice. |
| 02 | [Slices Internals and Usage](./02-slices-internals-and-usage.md) | The slice header (pointer, length, capacity), how `append` grows a slice, and slicing/aliasing semantics. |
| 03 | [Maps](./03-maps.md) | Hash maps, the nil-map trap, the comma-ok idiom for existence checks, and why iteration order is intentionally randomized. |
| 04 | [Structs and Composition](./04-structs-and-composition.md) | Struct literals and fields, and composition-over-inheritance via embedded structs. |
| 05 | [Pointers and Memory](./05-pointers-and-memory.md) | Pointer semantics with `&`/`*`, and a plain-language intro to when Go allocates on the stack vs. the heap. |
| 06 | [Methods and Value Vs Pointer Receivers](./06-methods-and-value-vs-pointer-receivers.md) | Declaring methods on types, and the rules for choosing a value receiver vs. a pointer receiver. |
| 07 | [Interfaces and Polymorphism](./07-interfaces-and-polymorphism.md) | Go's implicit interface satisfaction and the idiom of small, single-method interfaces. |
| 08 | [Type Assertions and Type Switches](./08-type-assertions-and-type-switches.md) | Extracting a concrete type from an interface value with `x.(T)` and `switch x.(type)`. |
| 09 | [Empty Interface and Reflection Intro](./09-empty-interface-and-reflection-intro.md) | The empty interface (`any`), and a first look at the `reflect` package — and when to avoid reaching for it. |

**Previous module:** [02 — Language Basics](../02-basics/README.md) · **Next module:** [04 — Concurrency](../04-concurrency/README.md)
