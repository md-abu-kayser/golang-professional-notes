# Programming Basics

## Overview

This lesson is a language-agnostic refresher on the concepts every programming language — including Go — is built from: variables, control flow, functions, and algorithmic thinking. If any of this is unfamiliar, spend real time here; every later module assumes it.

## Why It Matters

Go's syntax is deliberately small, which means you learn it fast — but only if the underlying concepts (state, iteration, decomposition) are already automatic. Skipping this step is the #1 reason beginners feel like Go is "hard" when the language itself is one of the simplest in wide use.

## Core Concepts

### Variables — named, typed storage

A variable is a name bound to a value in memory. In statically typed languages like Go, that binding also fixes a **type**, which determines what operations are valid and how much memory is reserved.

```text
age = 27          # a name bound to a value
age = age + 1     # the name can be rebound; the old value is discarded
```

### Control flow — deciding what runs

Programs branch (`if`/`else`) and repeat (loops) based on conditions:

```text
if temperature > 30:
    print("hot")
else:
    print("mild")

for i in 0..5:
    print(i)        # prints 0 1 2 3 4
```

Every loop needs three things: a starting state, a condition that eventually becomes false, and a step that moves toward that condition. Forgetting the third one is how you get an infinite loop.

### Functions — named, reusable behavior

A function packages a sequence of steps behind a name, optionally taking **parameters** in and returning a **value** out:

```text
function add(a, b):
    return a + b

result = add(2, 3)   # result = 5
```

Functions are the unit of decomposition: instead of one long script, you build a program out of small, named, testable pieces.

### Algorithmic thinking

An algorithm is just a precise, ordered set of steps to solve a problem. Before writing code, get in the habit of stating the steps in plain language first:

1. What's the input, and what shape is it in?
2. What's the desired output?
3. What's the smallest correct step-by-step path from one to the other?
4. What are the edge cases (empty input, one item, a huge input)?

This habit transfers directly to Go, where clarity of intent matters more than clever one-liners.

## Common Pitfalls

- **Confusing assignment with comparison** (`=` vs `==`) — Go's compiler catches this in most contexts, but the conceptual mistake still trips up the logic.
- **Off-by-one errors** in loops — always double-check whether a bound is inclusive or exclusive.
- **Writing code before stating the algorithm in words** — leads to solutions that "sort of" work but fall apart on edge cases.

## Key Takeaways

- Variables bind names to typed values; types constrain what you can do with them.
- Control flow (branching + looping) is how programs make decisions and repeat work.
- Functions are how you decompose a problem into named, reusable, testable pieces.
- Always state an algorithm in plain language before writing code — it's faster to debug English than syntax.
