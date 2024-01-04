package main

import "log"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 1; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 100; i++ {
		log.Printf("%d", <-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

/*
We can create worker jobs to always send and receive messages over channels
Here we are running 4 go routines that run a CPU intensive task of getting
nth fibonacci number.

This utilizes all the cores of the CPU and the CPU utilization jumps to nearly
320% as all 4 cores of my machine are being engaged in running worker-jobs.

*/
