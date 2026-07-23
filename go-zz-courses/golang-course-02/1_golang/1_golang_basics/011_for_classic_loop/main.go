package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// Sum numbers from 1 to N-1
	n := 10
	sum := 0

	for i := 0; i < n; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)
}
