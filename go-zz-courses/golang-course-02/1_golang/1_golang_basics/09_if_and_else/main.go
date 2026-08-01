package main

import "fmt"

func main() {
	purchaseAmount := 72000

	// if - else if - else
	if purchaseAmount >= 100000 {
		fmt.Println("Membership: Platinum")
	} else if purchaseAmount >= 75000 {
		fmt.Println("Membership: Gold")
	} else if purchaseAmount >= 45000 {
		fmt.Println("Membership: Silver")
	} else {
		fmt.Println("Membership: Bronze")
	}

	// Short statement with if
	if isEligible := purchaseAmount >= 45000; isEligible {
		fmt.Println("Is eligible for Exclusive Sale: ", isEligible)
	} else {
		fmt.Println("Is eligible for Exclusive Sale: ", isEligible)
	}
}

/*
Membership: Silver
Is eligible for Exclusive Sale:  true
*/
