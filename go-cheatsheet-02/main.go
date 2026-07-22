package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"
)

// ================================
// Go Cheat Sheet #2 (Intermediate)
// ================================

type Animal interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string { return "Woof" }

type Box[T any] struct {
	Value T
}

func printTitle(s string) {
	fmt.Printf("\n===== %s =====\n", s)
}

func main() {
	printTitle("Go Cheat Sheet #2")

	printTitle("Packages & Imports")
	fmt.Println("fmt, os, time, strings, sort, encoding/json")

	printTitle("Exported vs Unexported")
	fmt.Println("Capitalized identifiers are exported.")

	printTitle("Interfaces")
	var a Animal = Dog{}
	fmt.Println(a.Speak())

	printTitle("Type Assertion")
	var x any = "golang"
	if s, ok := x.(string); ok {
		fmt.Println(s)
	}

	printTitle("Type Switch")
	var v any = 42
	switch t := v.(type) {
	case int:
		fmt.Println("int:", t)
	case string:
		fmt.Println("string:", t)
	default:
		fmt.Println("unknown")
	}

	printTitle("Generics")
	b := Box[int]{Value: 100}
	fmt.Println(b.Value)

	printTitle("Reflection")
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(a))

	printTitle("Strings Builder")
	var sb strings.Builder
	sb.WriteString("Hello ")
	sb.WriteString("Go")
	fmt.Println(sb.String())

	printTitle("Time")
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Add(24 * time.Hour))

	printTitle("Sorting")
	nums := []int{5, 2, 8, 1}
	sort.Ints(nums)
	fmt.Println(nums)

	printTitle("JSON")
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	u := User{"Mohan", 30}
	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(data))

	var u2 User
	json.Unmarshal(data, &u2)
	fmt.Println(u2)

	printTitle("File I/O")
	os.WriteFile("sample.txt", []byte("Hello Go"), 0644)
	content, _ := os.ReadFile("sample.txt")
	fmt.Println(string(content))
	os.Remove("sample.txt")

	printTitle("Slice Tricks")
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s)

	printTitle("Map Idioms")
	m := map[string]int{"Go": 1}
	if val, ok := m["Go"]; ok {
		fmt.Println(val)
	}

	printTitle("Done")
}
