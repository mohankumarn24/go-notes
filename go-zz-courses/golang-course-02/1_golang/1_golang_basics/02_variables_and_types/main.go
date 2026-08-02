package main

import "fmt"

func main() {
	// var variableName type = value
	var year int = 1975
	fmt.Println("Year:", year) // Year: 1975

	var jacketPrice float64 = 24999.99
	fmt.Println("Jacket Price:", jacketPrice) // Jacket Price: 24999.99

	// Declaration without initialization
	// Zero value is assigned automatically.
	var brandName string
	brandName = "Emporio Armani"
	fmt.Println("Brand Name:", brandName) // Brand Name: Emporio Armani
}

/*
// Explicit declaration
var year int = 2026

// Explicit declaration
var year int
year = 2026

// Type inference using var
var year = 2026

// Short variable declaration (:=). Can only be used inside functions
year := 2026
*/
