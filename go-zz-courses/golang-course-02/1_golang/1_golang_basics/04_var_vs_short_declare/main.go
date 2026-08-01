package main

import "fmt"

func main() {
	// Explicit declaration
	var city string
	city = "Milan"

	// Type inference using var
	var brand = "Emporio Armani" // inferred as string

	// Short variable declaration (:=)
	// Can only be used inside functions.
	orders := 5000
	orders += 1000

	// Multiple short variable declarations
	jackets, shoes := 100, 30

	fmt.Println(city, brand, orders, jackets, shoes) // Milan Emporio Armani 6000 100 30
}
