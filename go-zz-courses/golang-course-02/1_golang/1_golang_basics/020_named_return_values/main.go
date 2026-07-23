package main

import "fmt"

// Function with named return values
func divide(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b

	return
}

func main() {
	quotient, remainder := divide(10, 3)

	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}
