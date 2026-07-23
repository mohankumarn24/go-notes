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
		Age:  20,
	}

	fmt.Println(u.Intro())

	u.Birthday() // Go automatically takes the address (&u)

	fmt.Println(u.Intro())
}

// Intro is a method on User.
//
// Value receiver means this method receives a copy of the User.
// Use a value receiver when the method doesn't modify the struct.
func (u User) Intro() string {
	return fmt.Sprintf("Hi, I am %s and I am %d years old.", u.Name, u.Age)
}

// Birthday is a method on User.
//
// Pointer receiver means this method receives the address of the User.
// Use a pointer receiver when the method needs to modify the struct.
func (u *User) Birthday() {
	u.Age++
}
