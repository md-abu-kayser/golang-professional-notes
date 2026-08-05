# First Program: Hello World

## Overview

Every Go program starts from the same three pieces: a `package` declaration, an `import` block, and (for executables) a `func main()`. This lesson writes and runs the smallest possible Go program and explains every line.

## Why It Matters

These three pieces appear in literally every Go file you will ever write or read. Understanding them cold now removes friction from every later lesson.

## Core Concepts

### The code

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Line by line

- **`package main`** — declares which package this file belongs to. `main` is special: it marks an executable program (as opposed to a reusable library package).
- **`import "fmt"`** — pulls in the standard library's formatting package, which provides `Println`, `Printf`, `Sprintf`, etc.
- **`func main()`** — the entry point. When you run a `main` package, execution starts here — no arguments, no return value.
- **`fmt.Println(...)`** — prints its arguments followed by a newline.

### Running it

Save the file as `main.go`, then:

```bash
go run main.go
# Hello, World!
```

`go run` compiles the program to a temporary binary and executes it in one step — ideal while developing. (See [04-go-commands](./04-go-commands-run-build-install-fmt.md) for `build` vs `run` vs `install`.)

## Common Pitfalls

- Naming the file anything and expecting `go run .` to "just work" without a `package main` — non-`main` packages can't be run directly.
- Forgetting the import — Go's compiler will refuse to build with an unused import *or* a missing one; there's no silent fallback.
- Expecting semicolons — Go inserts them automatically at compile time based on line breaks; writing them manually is legal but never idiomatic.

## Key Takeaways

- Every executable Go file needs `package main` and a `func main()`.
- `import` brings in packages you use — Go errors on both unused and missing imports.
- `go run file.go` compiles and executes in one step, ideal for quick iteration.
- `fmt.Println` is the simplest way to print output; `fmt.Printf` (covered later) adds formatting verbs.
