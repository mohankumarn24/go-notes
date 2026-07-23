package main

import "fmt"

// Variadic function
func sumAll(nums ...int) int {
	total := 0

	for _, currentValue := range nums {
		total += currentValue
	}

	return total
}

func main() {
	// Pass individual values
	fmt.Println("Sum:", sumAll(1, 2, 3, 4, 5))

	// Pass a slice using ... (slice expansion)
	values := []int{10, 23}
	fmt.Println("Sum:", sumAll(values...))

	// Anonymous function
	double := func(n int) int {
		return n * 2
	}

	fmt.Println("Double:", double(2))

	// IIFE (Immediately Invoked Function Expression)
	result := func(a, b int) int {
		return a + b
	}(5, 10)

	fmt.Println("Result:", result)
}
