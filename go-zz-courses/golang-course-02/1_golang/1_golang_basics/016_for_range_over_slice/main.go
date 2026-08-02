package main

import "fmt"

func main() {
	views := []int{10, 20, 30, 40, 50}

	// Iterate over a slice using range.
	// range returns index and value.
	total := 0
	for i, v := range views { // index, value -> i, v
		fmt.Println("Day:", i, "Views:", v)
		total += v
	}
	fmt.Println("Total Views:", total) // 150

	// Ignore the index using the blank identifier (_)
	for _, v := range views {
		fmt.Println("Views:", v)
	}
}
