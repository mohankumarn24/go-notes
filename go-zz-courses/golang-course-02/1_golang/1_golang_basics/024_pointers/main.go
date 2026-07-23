package main

import "fmt"

func main() {
	// Pointers store the memory address of a value.
	//
	// &x -> address of x (creates a pointer)
	// *p -> dereference (read or modify the value at that address)

	score := 10

	fmt.Println("Before:", score)
	fmt.Println("Address:", &score)

	addScore(&score)

	fmt.Println("After:", score)
}

// score is a pointer to an int.
func addScore(score *int) {
	// Dereference the pointer and modify the original value.
	// *score = *score + 5
	*score += 5
}
