func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs { results <- j*2 }
}