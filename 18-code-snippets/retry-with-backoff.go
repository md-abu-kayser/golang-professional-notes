package main

import (
	"fmt"
	"time"
)

func retryWithBackoff(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(1<<i) * time.Second)
	}
	return fmt.Errorf("after %d attempts: %w", attempts, err)
}

func main() {}
