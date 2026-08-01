package main

import "fmt"

func main() {
	jackets := 2
	pricePerJacket := 60000

	// Short statement with if
	if total := jackets * pricePerJacket; total >= 100000 {
		fmt.Println("Eligible for VIP Discount")
	} else {
		fmt.Println("Not Eligible for VIP Discount")
	}
}

// Eligible for VIP Discount
