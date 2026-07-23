package main

import "fmt"

func main() {
	views := []int{10, 20, 45, 50, 60}

	// Iterate over a slice using range.
	// range returns index and value.
	total := 0

	for i, v := range views {
		fmt.Println("Day:", i, "Views:", v)
		total += v
	}

	fmt.Println("Total Views:", total)

	// Ignore the index using the blank identifier (_)
	for _, v := range views {
		fmt.Println("Views:", v)
	}
}
