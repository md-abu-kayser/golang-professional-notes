# Editor Setup: VS Code + Go Extensions

## Overview

Go's tooling is designed to integrate deeply with editors via `gopls`, the official Go language server. This lesson sets up VS Code — the most common free choice — for a productive Go workflow; GoLand and Vim-go are noted as alternatives.

## Why It Matters

A correctly configured editor gives you real-time type errors, jump-to-definition, auto-imports, and formatting-on-save — turning the feedback loop from "compile and see" into "instant." This matters more in Go than in many languages because so much of Go's tooling (`vet`, `fmt`, test coverage) is designed to surface inline.

## Core Concepts

### Install the Go extension

1. Install [VS Code](https://code.visualstudio.com/).
2. Install the official **Go extension** (`golang.go`) from the marketplace.
3. Open a `.go` file — VS Code will prompt to install supporting tools (`gopls`, `dlv`, `staticcheck`, etc.). Accept this; it runs `go install` for each under the hood.

### What `gopls` gives you

`gopls` (the Go language server) powers:
- Autocomplete and inline type information
- Jump-to-definition / find-all-references
- Real-time error and warning squiggles
- Auto-import on save
- Rename-symbol refactoring across the whole module

### Recommended `settings.json`

```json
{
  "editor.formatOnSave": true,
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.testFlags": ["-v"],
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
  }
}
```

### Alternatives worth knowing

- **GoLand** (JetBrains) — a paid, full-featured Go IDE with deeper refactoring tools; a common choice on larger teams.
- **Vim-go / Neovim + `nvim-lspconfig`** — for terminal-first workflows, using the same `gopls` under the hood.

Whichever editor you choose, the underlying tool is the same `gopls` — the editor is just the interface to it.

## Common Pitfalls

- Skipping the "install supporting tools" prompt and then wondering why autocomplete doesn't work.
- Disabling format-on-save — leads to `gofmt` failures in CI that could've been caught locally.
- Running an old `gopls` version after a Go upgrade — if things feel broken after updating Go, run `go install golang.org/x/tools/gopls@latest`.

## Key Takeaways

- `gopls` is the language server behind Go tooling in every major editor — VS Code, GoLand, and Vim-go all use it.
- Enable format-on-save and organize-imports-on-save — it keeps every commit clean automatically.
- If autocomplete/errors aren't showing, the supporting tools (`gopls`, `dlv`, linters) likely didn't install — check the Go extension's output panel.
- Editor choice is personal; the tooling underneath (and therefore the behavior) is consistent across all of them.
