package main

import "fmt"

func main() {
	day := 3

	// Basic switch statement
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")
	}

	// Multiple values in a case
	grade := "A"

	switch grade {
	case "A", "A+":
		fmt.Println("Excellent")
	case "B", "B+":
		fmt.Println("Good")
	default:
		fmt.Println("Needs Improvement")
	}

	// Expression-less switch (alternative to if-else if)
	score := 82

	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 75:
		fmt.Println("Grade B")
	case score >= 45:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade D")
	}

	// fallthrough (rarely used)
	switch 1 {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
}
