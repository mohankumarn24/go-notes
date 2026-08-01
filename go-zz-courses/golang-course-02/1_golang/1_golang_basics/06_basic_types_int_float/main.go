package main

import "fmt"

func main() {
	// Integer arithmetic
	views1 := 1000
	views2 := 2000
	totalViews := views1 + views2
	fmt.Println("Total Views:", totalViews) // 3000

	likes := 10
	likes++
	fmt.Println("Likes:", likes) // 11

	// Integer division
	avgViews := totalViews / 2
	fmt.Println("Average Views:", avgViews) // 1500

	// Floating-point arithmetic
	rating1 := 4.5
	rating2 := 5.1
	avgRating := (rating1 + rating2) / 2
	fmt.Println("Average Rating:", avgRating) // 4.8

	// Arithmetic operators
	a, b := 7, 2

	fmt.Println("Addition:", a+b)       // 9
	fmt.Println("Subtraction:", a-b)    // 5
	fmt.Println("Multiplication:", a*b) // 14
	fmt.Println("Division:", a/b)       // Integer division: 3
	fmt.Println("Modulus:", a%b)        // 1
}

/*
Total Views: 3000
Likes: 12
Average Views: 1500
Average Rating: 4.8
Addition: 13
Subtraction: 7
Multiplication: 30
Division: 3
Modulus: 1
*/
