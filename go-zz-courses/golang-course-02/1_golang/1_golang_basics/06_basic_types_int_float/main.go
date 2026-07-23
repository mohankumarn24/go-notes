package main

import "fmt"

func main() {
	// Integer arithmetic
	views1 := 1000
	views2 := 2000
	totalViews := views1 + views2

	likes := 10
	likes++
	likes++

	avgViews := totalViews / 2 // Integer division

	fmt.Println("Total Views:", totalViews)
	fmt.Println("Likes:", likes)
	fmt.Println("Average Views:", avgViews)

	// Floating-point arithmetic
	rating1 := 4.5
	rating2 := 5.1

	avgRating := (rating1 + rating2) / 2

	fmt.Println("Average Rating:", avgRating)

	// Arithmetic operators
	a, b := 10, 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b) // Integer division
	fmt.Println("Modulus:", a%b)
}
