# Command Line Basics

## Overview

Go's tooling (`go build`, `go run`, `go test`, module management) is entirely command-line driven — there's no IDE-only workflow. This lesson covers the terminal fluency you need before installing Go itself.

## Why It Matters

Every Go tutorial, every CI pipeline, and every production deployment assumes shell comfort. Fumbling with navigation slows down every single lesson after this one.

## Core Concepts

### Navigating the filesystem

```bash
pwd                 # print current directory
ls -la               # list contents, including hidden files, with details
cd path/to/dir        # change directory
cd ..                 # move up one level
cd ~                  # jump to home directory
```

### Working with files and directories

```bash
mkdir my-project      # create a directory
touch main.go          # create an empty file
cat main.go             # print file contents
rm main.go               # delete a file
rm -rf my-project         # delete a directory and its contents (be careful!)
cp source.go dest.go       # copy a file
mv old-name.go new-name.go  # rename or move a file
```

### Environment variables and `PATH`

The shell keeps a set of key-value variables available to every program you run. The most important one for Go is `PATH` — the list of directories the shell searches when you type a command name.

```bash
echo $PATH                       # show the current PATH
export PATH=$PATH:/usr/local/go/bin  # append Go's bin directory (Linux/macOS)
```

When `go: command not found` shows up after installing Go, it's almost always a `PATH` issue — the installer put the binary somewhere the shell isn't looking.

### Redirection and piping

```bash
go build . > build.log 2>&1   # send both stdout and stderr to a file
cat build.log | grep error     # pipe output from one command into another
```

## Common Pitfalls

- Running commands from the wrong directory — always check `pwd` when a command "isn't working."
- Using `rm -rf` carelessly — there's no trash can on the command line.
- Forgetting that `PATH` changes made in one terminal session don't persist — they need to go in your shell's config file (`.bashrc`, `.zshrc`) to survive a restart.

## Key Takeaways

- `pwd`, `ls`, `cd` cover 90% of daily navigation.
- `PATH` is why a command works in one terminal and not another — check it first when a tool "isn't found."
- Redirection (`>`, `>>`) and piping (`|`) let you chain commands instead of copy-pasting output by hand.
- Shell config changes (`PATH`, aliases) belong in `.bashrc`/`.zshrc`/`.zprofile`, not just typed ad hoc.
