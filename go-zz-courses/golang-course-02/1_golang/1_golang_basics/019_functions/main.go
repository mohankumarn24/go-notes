package main

import "fmt"

// Function with parameters and a single return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func sumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}

// Function with named return values
func divideAndRemainder(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// Variadic function
func total(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	s, p := sumAndProduct(6, 5)
	fmt.Println("Sum:", s)
	fmt.Println("Product:", p)

	// Ignore unwanted return values using the blank identifier (_)
	onlySum, _ := sumAndProduct(10, 2)
	fmt.Println("Only Sum:", onlySum)

	q, r := divideAndRemainder(10, 3)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)

	fmt.Println("Total:", total(1, 2, 3, 4, 5))

	// Anonymous function
	square := func(n int) int {
		return n * n
	}

	fmt.Println("Square:", square(5))
}
