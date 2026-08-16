# Contributing to Golang Professional Notes

Thanks for considering a contribution — this repo grows faster and stays more accurate with outside eyes on it.

## Ways to contribute

- **Fill in a stub lesson.** Every planned (📋) or syllabus-ready (🚧) file already has a title and learning objectives — turn it into a full lesson.
- **Fix an error** in an existing note or code snippet.
- **Improve an example** — a clearer explanation, a better diagram, a more realistic code sample.
- **Report a gap** via an issue if something's missing or out of order.

## Workflow

1. **Fork** the repository and create a branch off `main`:
   ```bash
   git checkout -b lesson/04-concurrency-goroutines
   ```
2. **Make your change**, following the lesson template below.
3. **Run formatting checks** on any Go code:
   ```bash
   gofmt -l .
   go vet ./...
   ```
4. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/):
   - `docs: add goroutines lesson (04-concurrency/01)`
   - `fix: correct buffered channel example in 02-basics`
   - `feat: add worker-pool code snippet`
5. **Open a pull request** against `main` using the PR template — describe what changed and why.

## Lesson template

Every lesson file follows this structure so the repo reads as one consistent book:

```markdown
# Lesson Title

## Overview

One paragraph: what this concept is and where it fits.

## Why It Matters

Real-world motivation — what breaks or gets harder without this concept.

## Core Concepts

The explanation, with idiomatic, runnable Go code blocks.

## Common Pitfalls

Mistakes learners actually make with this topic.

## Key Takeaways

3–5 bullet points, the "if you remember nothing else" summary.
```

## Style guide

- Code must be `gofmt`-formatted and pass `go vet`.
- Prefer standard library examples before introducing third-party packages.
- Every code block should be runnable as-is or clearly marked as pseudocode.
- Keep prose direct — this is a reference, not a narrative.

## Code of conduct

Be respectful, assume good faith, and keep feedback focused on the work. Disagreements about content are welcome; personal attacks are not.
