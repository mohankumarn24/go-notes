package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// sync.Once
	// ---------
	// Ensures a function is executed exactly once,
	// even if multiple goroutines call it simultaneously.
	//
	// Common use cases:
	// 1. Initialize a database connection.
	// 2. Load configuration.
	// 3. Create a Singleton instance.
	// 4. Initialize a logger or cache.

	var wg sync.WaitGroup
	var once sync.Once

	workers := 5

	// Wait for all workers to finish.
	wg.Add(workers)

	for i := 1; i <= workers; i++ {

		go func(id int) {
			defer wg.Done()

			// Only one goroutine executes setup().
			// All others wait until it completes,
			// then continue without calling it again.
			once.Do(setup)

			fmt.Printf("Worker %d: running\n", id)

		}(i)
	}

	wg.Wait()

	fmt.Println("Main: all workers finished")
}

// Runs exactly once.
func setup() {

	fmt.Println("Setup: initialization started")

	// Simulate an expensive initialization.
	time.Sleep(600 * time.Millisecond)

	fmt.Println("Setup: initialization completed")
}

/*
Setup: initialization started
Setup: initialization completed
Worker 1: running
Worker 5: running
Worker 2: running
Worker 3: running
Worker 4: running
Main: all workers finished
*/
