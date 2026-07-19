package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error example.
var ErrNotFound = errors.New("not found")

func doSomething(ok bool) error {
	if !ok {
		return fmt.Errorf("doSomething failed: %w", ErrNotFound)
	}
	return nil
}

func main() {
	if err := doSomething(false); err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("not found:", err)
		}
	}
}
