# What Is Go?

## Overview

Go (often called Golang) is an open-source, compiled, statically typed language created at Google and released publicly in 2009. It was designed by Robert Griesemer, Rob Pike, and Ken Thompson to solve a specific, practical problem: building reliable software fast, at scale, with large teams — not to explore new language theory.

## Why It Matters

Understanding _why_ Go looks the way it does explains almost every design decision you'll hit later: why there's no exceptions, why there's no generics-heavy stdlib, why formatting isn't a matter of taste (`gofmt` decides), and why concurrency is a first-class language feature instead of a library bolted on.

## Core Concepts

### The problem Go was built to solve

At Google's scale, the pain points weren't "we need a clever type system" — they were: slow builds on huge codebases, difficulty reasoning about concurrent code, and engineers spending time on formatting debates and dependency hell instead of shipping. Go's answer:

- **Fast compilation** — even large codebases build in seconds.
- **Simplicity over expressiveness** — a small language spec (readable in an afternoon) that's easy to onboard new engineers into and easy to read in code review.
- **Built-in concurrency** — goroutines and channels as language primitives, not a threading library.
- **One way to format code** — `gofmt` ends style debates entirely.
- **Static binaries** — compile to a single executable with no runtime dependency to install.

### Where Go is used in production

Go is the language behind much of modern cloud infrastructure: Docker, Kubernetes, Terraform, and Prometheus are all written in Go. It's a common choice for backend APIs, CLIs, network services, and anywhere high concurrency and predictable performance matter more than a large ecosystem of language features.

### What Go deliberately leaves out

- No exceptions — errors are explicit return values (see [02-basics/07](../02-basics/07-errors-and-error-handling-pattern.md)).
- No classical inheritance — composition via embedded structs instead (see [03-intermediate-types/04](../03-intermediate-types/04-structs-and-composition.md)).
- No implicit type conversions — you convert explicitly, on purpose.
- (Until Go 1.18) no generics — added later, deliberately, once a design that fit Go's philosophy was found (see [16-advanced-concepts/01](../16-advanced-concepts/01-generics-and-type-parameters.md)).

## Common Pitfalls

- Judging Go by feature-count against languages like Rust or C++ — Go optimizes for a different variable: engineering velocity across large teams, not language expressiveness.
- Assuming "simple language" means "simple problems" — Go runs some of the highest-scale distributed systems in the world.

## Key Takeaways

- Go was built at Google to solve real engineering-at-scale problems: build speed, concurrency, team onboarding.
- Simplicity is a deliberate design goal, not a limitation — a small spec is a feature.
- Concurrency (goroutines, channels) is built into the language itself.
- Go underpins much of the modern cloud-native ecosystem (Docker, Kubernetes, Terraform).
