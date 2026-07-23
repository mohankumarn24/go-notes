package main

import "fmt"

// User represents a person.
type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "Sangam",
		Age:  13,
	}

	fmt.Println("Before Birthday:", u.Age)

	// Go automatically takes the address (&u)
	// because Birthday has a pointer receiver.
	u.Birthday()

	fmt.Println("After Birthday:", u.Age)
}

// Pointer receiver means this method receives
// the address of the User.
//
// Use a pointer receiver when the method
// needs to modify the struct.
func (u *User) Birthday() {
	u.Age++
}
