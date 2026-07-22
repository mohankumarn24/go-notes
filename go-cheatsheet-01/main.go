package main

import (
	"fmt"
	"strings"
	"time"
)

// ==========================
// Go Cheat Sheet #1 (Basics)
// ==========================

const Pi = 3.14159

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, I'm %s (%d)\n", p.Name, p.Age)
}

func add(a, b int) int { return a + b }

func swap(a, b int) (int, int) { return b, a }

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	panic("example panic")
}

func main() {
	fmt.Println("=== Go Cheat Sheet #1 ===")

	// Variables
	var x int = 10
	y := 20
	var z int
	fmt.Println(x, y, z)

	// Constants
	fmt.Println("Pi:", Pi)

	// Types
	var (
		i  int     = 42
		f  float64 = 3.14
		s  string  = "Go"
		b  bool    = true
		r  rune    = 'A'
		by byte    = 'Z'
	)
	fmt.Println(i, f, s, b, r, by)

	// Arrays
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// Slices
	sl := []string{"Java", "Go"}
	sl = append(sl, "Python")
	fmt.Println(sl)

	// Maps
	m := map[string]int{"Alice": 90, "Bob": 80}
	m["Carol"] = 95
	fmt.Println(m)

	// If
	if x < y {
		fmt.Println("x<y")
	}

	// Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// Loops
	for i := 1; i <= 3; i++ {
		fmt.Println("for", i)
	}

	n := 0
	for n < 3 {
		fmt.Println("while", n)
		n++
	}

	for idx, val := range sl {
		fmt.Println(idx, val)
	}

	// Break/Continue
	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// Functions
	fmt.Println(add(5, 7))
	a, b2 := swap(1, 2)
	fmt.Println(a, b2)
	fmt.Println(sum(1, 2, 3, 4))

	// Anonymous function
	square := func(v int) int { return v * v }
	fmt.Println(square(5))

	// Closure
	counter := func() func() int {
		c := 0
		return func() int {
			c++
			return c
		}
	}()
	fmt.Println(counter(), counter())

	// Pointers
	p := &x
	*p = 100
	fmt.Println(x)

	// Structs & Methods
	person := Person{"Mohan", 30}
	person.Greet()

	// Strings
	fmt.Println(strings.ToUpper("golang"))
	fmt.Println(strings.Contains("golang", "go"))

	// Type conversion
	fmt.Println(float64(x) / 3)

	// Error handling
	if q, err := divide(10, 2); err == nil {
		fmt.Println(q)
	}
	if _, err := divide(10, 0); err != nil {
		fmt.Println(err)
	}

	// defer
	defer fmt.Println("Deferred execution")

	// panic / recover
	recoverDemo()

	fmt.Println("End")
}
