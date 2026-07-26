package main

import (
	"fmt"
	"time"
)

func main() {

	// Buffered Channel
	// ----------------
	// Capacity = 2
	// The producer can send up to 2 values without blocking.
	// The 3rd send blocks until the consumer receives one value.

	jobs := make(chan string, 2)

	// Producer Goroutine
	go func() {

		fmt.Println("Producer: sending job-1")
		jobs <- "job-1" // Buffer: [job-1]

		fmt.Println("Producer: sending job-2")
		jobs <- "job-2" // Buffer: [job-1, job-2] (Full)

		fmt.Println("Producer: sending job-3 (waiting if buffer is full)")
		jobs <- "job-3" // Blocks until consumer removes one job

		fmt.Println("Producer: all jobs sent")

		// No more values will be sent.
		close(jobs)
	}()

	// Consumer (Main Goroutine)
	// range automatically keeps receiving until:
	// 1. Channel is closed
	// 2. Buffer becomes empty
	for job := range jobs {

		fmt.Println("Consumer: received", job)

		// Simulate processing time
		time.Sleep(300 * time.Millisecond)

		fmt.Println("Consumer: finished", job)
	}

	fmt.Println("Main: all jobs completed")
}

/*
Producer: sending job-1
Producer: sending job-2
Producer: sending job-3 (waiting if buffer is full)
Producer: all jobs sent
Consumer: received job-1
Consumer: finished job-1
Consumer: received job-2
Consumer: finished job-2
Consumer: received job-3
Consumer: finished job-3
Main: all jobs completed
*/
