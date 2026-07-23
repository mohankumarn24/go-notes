package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0, // Valid value
	}

	fmt.Println("a:", points["a"])
	fmt.Println("b:", points["b"])
	fmt.Println("c:", points["c"]) // Missing key returns the zero value

	// Check if a key exists (comma-ok idiom)
	valB, okB := points["b"]
	fmt.Println("b:", valB, "Exists:", okB)

	valC, okC := points["c"]
	fmt.Println("c:", valC, "Exists:", okC)

	// Short statement with if
	if val, ok := points["b"]; ok {
		fmt.Println("b is present:", val)
	} else {
		fmt.Println("b key is not present")
	}

	prices := map[string]int{
		"xyz": 500,
		"def": 1800,
	}

	// Iterate over a map (key-value pairs)
	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total += price
	}
	fmt.Println("Total:", total)

	// Iterate over keys only
	for item := range prices {
		fmt.Println("Item:", item)
	}

	// Iterate over values only
	for _, price := range prices {
		fmt.Println("Price:", price)
	}
}
