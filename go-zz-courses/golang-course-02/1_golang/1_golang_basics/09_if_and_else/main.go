package main

import "fmt"

func main() {
	score := 72

	// if - else if - else
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 45 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// Short statement with if
	if passed := score >= 45; passed {
		fmt.Println("Result: Pass")
	} else {
		fmt.Println("Result: Fail")
	}
}
