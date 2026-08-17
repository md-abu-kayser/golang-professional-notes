for i := 0; i < maxRetries; i++ {
    if err = doWork(); err == nil { break }
    time.Sleep(time.Duration(1<<i) * time.Second) // exponential backoff
}