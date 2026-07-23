package main

import "fmt"

// Struct groups related fields into one custom type.
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	// Create a struct using field names.
	u1 := User{
		ID:    1,
		Name:  "Sangam",
		Email: "sangam@gmail.com",
		Age:   100,
	}

	fmt.Println("User 1:", u1)
	fmt.Println("ID:", u1.ID)
	fmt.Println("Email:", u1.Email)

	// Struct fields are mutable by default.
	u1.Age = 200

	fmt.Println("Updated User 1:", u1)

	// Partial initialization.
	// Unspecified fields receive their zero values.
	u2 := User{
		Name:  "John",
		Email: "john@gmail.com",
	}

	fmt.Println("User 2:", u2)
}
