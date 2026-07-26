package main

import (
	"fmt"
	"time"
)

func main() {

	// select
	// ------
	// Waits on multiple channel operations.
	// Executes the first case that becomes ready.
	//
	// Common use cases:
	// 1. Wait for a worker result.
	// 2. Timeout if it takes too long.
	// 3. Handle cancellation (context.Done()).

	// Channel that carries the worker's result.
	resultCh := make(chan string)

	// Worker Goroutine
	go func() {

		// Simulate a slow operation
		// (Database call, HTTP request, Microservice call, etc.)
		time.Sleep(40 * time.Millisecond)

		resultCh <- "Worker: success"
	}()

	// Timeout channel.
	// After 250 ms, Go automatically sends a value on this channel.
	timeoutCh := time.After(250 * time.Millisecond)

	// Wait for whichever happens first.
	select {

	case result := <-resultCh:
		fmt.Println("Main: received result ->", result)

	case <-timeoutCh:
		fmt.Println("Main: timeout occurred. Stopped waiting.")
	}

	fmt.Println("Main: work completed")
}

/*
Main: received result -> Worker: success
Main: work completed
*/
