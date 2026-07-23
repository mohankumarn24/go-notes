package main

import "fmt"

func main() {
	// Explicit declaration
	var city string
	city = "London"

	// Type inference using var
	var channel = "Sangam" // inferred as string

	// Short variable declaration (:=)
	// Can only be used inside functions.
	subscribers := 5000
	subscribers += 1000

	// Multiple short variable declarations
	likes, comments := 100, 30

	fmt.Println(city, channel, subscribers, likes, comments)
}
