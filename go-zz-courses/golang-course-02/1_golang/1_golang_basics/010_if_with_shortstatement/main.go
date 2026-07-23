package main

import "fmt"

func main() {
	items := 3
	pricePerItem := 40

	// Short statement with if
	if total := items * pricePerItem; total >= 100 {
		fmt.Println("Eligible for Free Shipping")
	} else {
		fmt.Println("Not Eligible for Free Shipping")
	}
}
