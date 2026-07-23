package main

import "fmt"

func main() {
	// Constants cannot be changed after declaration.

	// Untyped constant
	const appName = "Go Basics"

	// Typed constants
	const maxUpload int = 25
	const discountedPrice float64 = 10.3
	const isProduction bool = false

	// Constant expression (evaluated at compile time)
	const daysInWeek = 7
	const weeks = 4
	const daysInMonth = daysInWeek * weeks

	fmt.Println("App Name:", appName)
	fmt.Println("Max Upload:", maxUpload)
	fmt.Println("Discounted Price:", discountedPrice)
	fmt.Println("Production:", isProduction)
	fmt.Println("Days in Month:", daysInMonth)
}
