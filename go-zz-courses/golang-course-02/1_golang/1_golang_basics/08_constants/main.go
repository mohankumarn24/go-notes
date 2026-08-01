package main

import "fmt"

func main() {
	// Constants must be initialized when declared
	// Constants cannot be modified later

	// Typed constants
	const maxUpload int = 25
	const discountedPrice float64 = 10.3
	const isProduction bool = false

	/*
		// missing init expr for year
		const year int
		year = 2026
	*/

	// Untyped constant
	const appName = "Go Basics"

	// Constant expressions (evaluated at compile time)
	const daysInWeek = 7
	const weeks = 4
	const daysInMonth = daysInWeek * weeks

	fmt.Println("App Name:", appName)
	fmt.Println("Max Upload:", maxUpload)
	fmt.Println("Discounted Price:", discountedPrice)
	fmt.Println("Production:", isProduction)
	fmt.Println("Days in Month:", daysInMonth)
}

/*
App Name: Go Basics
Max Upload: 25
Discounted Price: 10.3
Production: false
Days in Month: 28
*/
