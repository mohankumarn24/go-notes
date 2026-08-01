package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println("Traditional: ", i)
	}

	// Modern Go (Go 1.22+)
	// range over an integer generates values from 0 to n-1.
	for i := range 10 {
		fmt.Println("Range: ", i)
	}

	// Sum numbers from 0 to n-1.
	n := 10
	sum := 0
	for i := range n {
		sum += i
	}
	fmt.Println("Sum: ", sum)

	// Infinite loop
	count := 1
	for {
		fmt.Println("Count: ", count)

		if count == 5 {
			break // Exit the loop
		}

		count++
	}
	fmt.Println("Loop finished")

	// Go doesn't have a while keyword.
	// Use for as a while loop.
	j := 0

	for j < 10 {
		fmt.Println("While: ", j)
		j++
	}

	// Go doesn't have a do...while loop.
	// Use an infinite loop with break when the body
	// must execute at least once.
	doCount := 1

	for {
		fmt.Println("Do-While: ", doCount)
		doCount++

		if doCount > 5 {
			break
		}
	}
}

/*
Traditional:  0
Traditional:  1
Traditional:  2
Traditional:  3
Traditional:  4
Traditional:  5
Traditional:  6
Traditional:  7
Traditional:  8
Traditional:  9
Range:  0
Range:  1
Range:  2
Range:  3
Range:  4
Range:  5
Range:  6
Range:  7
Range:  8
Range:  9
Sum:  45
Count:  1
Count:  2
Count:  3
Count:  4
Count:  5
Loop finished
While:  0
While:  1
While:  2
While:  3
While:  4
While:  5
While:  6
While:  7
While:  8
While:  9
Do-While:  1
Do-While:  2
Do-While:  3
Do-While:  4
Do-While:  5
*/
