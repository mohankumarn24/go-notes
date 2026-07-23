package main

import (
	"errors"
	"fmt"
)

func main() {
	// Example: Failure path
	fmt.Println("Case 1: Fail Early")

	if err := doWork(false); err != nil {
		fmt.Println("Error:", err)
	}

	// Uncomment to test the success path.
	/*
		fmt.Println("\nCase 2: Success")

		if err := doWork(true); err != nil {
			fmt.Println("Error:", err)
		}
	*/
}

func doWork(success bool) error {
	// Simulate acquiring a resource.
	fmt.Println("Start: Resource Acquired")

	// Common uses of defer:
	// - defer file.Close()
	// - defer resp.Body.Close()
	// - defer mutex.Unlock()

	// defer guarantees this runs before the function returns,
	// whether it returns normally or returns early due to an error.
	defer fmt.Println("Cleanup: Resource Released")

	if !success {
		return errors.New("something went wrong; returning early")
	}

	fmt.Println("Work: Doing something important")
	fmt.Println("Work: Completed")

	return nil
}
