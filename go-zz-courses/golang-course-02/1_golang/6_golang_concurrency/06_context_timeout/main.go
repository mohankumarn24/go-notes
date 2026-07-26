package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Context
	// -------
	// Context is used to:
	// 1. Set timeouts.
	// 2. Cancel long-running work.
	// 3. Propagate cancellation across goroutines.

	// Create a context that automatically expires after 450 ms.
	ctx, cancel := context.WithTimeout(context.Background(), 450*time.Millisecond)

	// Always call cancel to release resources.
	defer cancel()

	// Start a slow worker.
	go slowWork(ctx)

	fmt.Println("Main: waiting for context to finish...")

	// Wait until the context is cancelled or times out.
	<-ctx.Done()

	fmt.Println("Main: context ended ->", ctx.Err())
	fmt.Println("Main: exiting")
}

func slowWork(ctx context.Context) {

	fmt.Println("Worker: started")

	select {

	// Simulate a slow operation (700 ms).
	case <-time.After(700 * time.Millisecond):
		fmt.Println("Worker: completed successfully")

	// Context timed out or was cancelled.
	case <-ctx.Done():
		fmt.Println("Worker: cancelled ->", ctx.Err())
	}
}

/*
Main: waiting for context to finish...
Worker: started
Main: context ended -> context deadline exceeded
Main: exiting
*/
