package main

import (
	"fmt"
	"time"
)

func main() {
	// Concurrency = multiple tasks make progress together.
	// Parallelism = multiple tasks execute at the same time.
	//
	// Goroutines provide concurrency. Whether they run in parallel
	// depends on the number of CPU cores available.

	start := time.Now()

	// Start two independent tasks.
	// The 'go' keyword starts a new goroutine and DOES NOT wait for it.
	go fetchUserService(start)
	go fetchOrderService(start)

	fmt.Println("Main: started both service calls at", time.Since(start))

	// Main goroutine continues doing its own work.
	fmt.Println("Main: processing request - Step 1", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Main: processing request - Step 2", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Main: processing request - Step 3", time.Since(start))

	// Demo only.
	// Without this sleep, main may exit before the goroutines finish.
	// Later we'll replace this with sync.WaitGroup.
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Main: exiting at", time.Since(start))
}

// Simulates calling a User microservice.
func fetchUserService(start time.Time) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("User Service: response received at", time.Since(start))
}

// Simulates calling an Order microservice.
func fetchOrderService(start time.Time) {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Order Service: response received at", time.Since(start))
}

/*
Output:
Main: started both service calls at 0s
Main: processing request - Step 1 0s
Main: processing request - Step 2 100.5872ms
Order Service: response received at 150.5592ms
Main: processing request - Step 3 201.2665ms
User Service: response received at 300.5843ms
Main: exiting at 701.6655ms
*/

/***********************************************
package main

import (
	"fmt"
	"time"
)

func main() {
	// main tarts a goroutine and does not wait for its completion
	go func() {
		// Temporarily delays the current goroutine (not a synchronization mechanism)
		time.Sleep(1 * time.Second)
		fmt.Println("Hello from goroutine")
	}()

	fmt.Println("Main finished")
}

Output: Main finished
***********************************************/

/***********************************************
var wg sync.WaitGroup

wg.Add(2)

go func() {
	defer wg.Done()
	fetchUserService()
}()

go func() {
	defer wg.Done()
	fetchOrderService()
}()

wg.Wait() // Wait until both goroutines call Done()

fmt.Println("All goroutines finished")
***********************************************/

/***********************************************
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	// We are waiting for 2 goroutines.
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done() // Notify WaitGroup when finished.
		fetchUserService(start)
	}()

	// Goroutine 2
	go func() {
		defer wg.Done() // Notify WaitGroup when finished.
		fetchOrderService(start)
	}()

	fmt.Println("Main: started both service calls at", time.Since(start))

	// Main can continue doing other work.
	fmt.Println("Main: processing request - Step 1", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Main: processing request - Step 2", time.Since(start))

	// Wait until both goroutines finish.
	wg.Wait()

	fmt.Println("Main: all goroutines finished at", time.Since(start))
	fmt.Println("Main: exiting")
}

func fetchUserService(start time.Time) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("User Service: response received at", time.Since(start))
}

func fetchOrderService(start time.Time) {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Order Service: response received at", time.Since(start))
}


Output:
Main: started both service calls at 18.2µs
Main: processing request - Step 1 42.6µs
Main: processing request - Step 2 100.4ms
Order Service: response received at 150.8ms
User Service: response received at 300.9ms
Main: all goroutines finished at 301.0ms
Main: exiting
***********************************************/
