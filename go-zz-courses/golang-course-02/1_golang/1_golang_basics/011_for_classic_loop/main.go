package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Traditional:", i)
	}

	// Modern Go (Go 1.22+)
	// range over an integer generates values from 0 to n-1.
	for i := range 5 {
		fmt.Println("Range:", i)
	}

	// Sum numbers from 0 to n-1.
	n := 10
	sum := 0

	for i := range n {
		sum += i
	}

	fmt.Println("Sum:", sum)

	// Infinite loop
	count := 1

	for {
		fmt.Println("Count:", count)

		if count == 5 {
			break // Exit the loop
		}

		count++
	}

	fmt.Println("Loop finished")

	// Go doesn't have a while keyword.
	// Use for as a while loop.
	x := 1

	for x <= 3 {
		fmt.Println("While:", x)
		x++
	}

	// Go doesn't have a do...while loop.
	// Use an infinite loop with break when the body
	// must execute at least once.
	doCount := 1

	for {
		fmt.Println("Do-While:", doCount)
		doCount++

		if doCount > 5 {
			break
		}
	}
}
