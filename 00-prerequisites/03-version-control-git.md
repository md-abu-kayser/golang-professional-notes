# Version Control with Git

## Overview

Git tracks changes to a codebase over time and enables collaboration without overwriting each other's work. Every Go project — from a weekend script to a production microservice — lives in a Git repository.

## Why It Matters

Without version control, "undo" only goes back one step and collaboration means emailing zip files. Git gives you a full history, safe experimentation via branches, and the collaboration model (pull requests, code review) used by essentially every professional engineering team.

## Core Concepts

### The basic loop: status → add → commit

```bash
git init                    # start tracking a new project
git status                   # see what's changed
git add main.go                # stage a specific file
git add .                       # stage everything changed
git commit -m "Add HTTP server"  # save a snapshot with a message
```

### Connecting to a remote (GitHub)

```bash
git remote add origin https://github.com/username/repo.git
git push -u origin main       # push local commits, set upstream
git pull                       # fetch + merge remote changes
git clone https://github.com/username/repo.git  # copy an existing repo
```

### Branching and merging

Branches let you develop a feature in isolation without touching working code:

```bash
git checkout -b feature/jwt-auth   # create and switch to a new branch
# ...make changes, commit...
git checkout main                    # switch back
git merge feature/jwt-auth            # bring the feature branch in
```

### `.gitignore`

Tells Git which files to never track — compiled binaries, secrets, editor config:

```text
bin/
*.env
.vscode/
```

### Commit message conventions

This repo (and most professional Go projects) uses [Conventional Commits](https://www.conventionalcommits.org/):

```text
feat: add JWT middleware
fix: correct nil pointer in worker pool shutdown
docs: expand goroutines lesson with select example
```

The prefix (`feat`, `fix`, `docs`, `refactor`, `test`) makes history scannable and can auto-generate changelogs.

## Common Pitfalls

- Committing secrets (`.env`, API keys) before adding them to `.gitignore` — once pushed, treat the secret as compromised and rotate it.
- Vague commit messages like `"fix stuff"` — useless when debugging six months later with `git blame`.
- Working directly on `main` instead of feature branches — makes code review and rollback much harder.

## Key Takeaways

- The daily loop is `status` → `add` → `commit`; `push`/`pull` sync with a remote.
- Branches isolate work in progress; `main` should stay deployable.
- `.gitignore` keeps binaries and secrets out of history — set it up before the first commit.
- Clear, conventional commit messages are documentation you write for future-you.
