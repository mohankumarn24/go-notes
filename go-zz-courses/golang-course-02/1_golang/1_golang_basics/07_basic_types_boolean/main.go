package main

import "fmt"

func main() {
	// Boolean variables
	isLogged := true
	isAdmin := true
	hasSubscription := false

	// Logical operators
	canOpenDashboard := isLogged && hasSubscription
	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println("Can Open Dashboard:", canOpenDashboard)
	fmt.Println("Can Delete Post:", canDeletePost)

	// Logical NOT
	fmt.Println("!isLogged:", !isLogged)

	// Comparison operators
	age := 20
	isAdult := age >= 18
	fmt.Println("Is Adult:", isAdult)

	fmt.Println("10 > 5 :", 10 > 5)
	fmt.Println("10 >= 10:", 10 >= 10)

	fmt.Println("10 < 5 :", 10 < 5)
	fmt.Println("10 <= 9 :", 10 <= 9)

	fmt.Println("10 == 10:", 10 == 10)
	fmt.Println("10 != 5 :", 10 != 5)
}
