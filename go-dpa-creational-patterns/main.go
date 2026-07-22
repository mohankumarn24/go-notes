package main

import "fmt"

// ================================
// Go Creational Design Patterns
// Single runnable program
// ================================

// ---------- Singleton ----------
type singleton struct{}

var instance = &singleton{}

func GetSingleton() *singleton { return instance }

// ---------- Factory ----------
type Shape interface{ Draw() }

type Circle struct{}

func (Circle) Draw() { fmt.Println("Drawing Circle") }

type Square struct{}

func (Square) Draw() { fmt.Println("Drawing Square") }

func NewShape(kind string) Shape {
	switch kind {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

// ---------- Abstract Factory ----------
type Button interface{ Click() }
type Checkbox interface{ Check() }

type WinButton struct{}

func (WinButton) Click() { fmt.Println("Windows Button") }

type WinCheckbox struct{}

func (WinCheckbox) Check() { fmt.Println("Windows Checkbox") }

type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (WindowsFactory) CreateButton() Button     { return WinButton{} }
func (WindowsFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

// ---------- Builder ----------
type House struct {
	Walls  int
	Doors  int
	Garden bool
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) Walls(n int) *HouseBuilder {
	b.house.Walls = n
	return b
}
func (b *HouseBuilder) Doors(n int) *HouseBuilder {
	b.house.Doors = n
	return b
}
func (b *HouseBuilder) Garden(v bool) *HouseBuilder {
	b.house.Garden = v
	return b
}
func (b *HouseBuilder) Build() House {
	return b.house
}

// ---------- Prototype ----------
type Document struct {
	Title string
	Pages int
}

func (d Document) Clone() Document {
	return d
}

// ---------- Demos ----------
func singletonDemo() {
	fmt.Println("Singleton")
	s1 := GetSingleton()
	s2 := GetSingleton()
	fmt.Println("Same instance?", s1 == s2)
}

func factoryDemo() {
	fmt.Println("Factory")
	NewShape("circle").Draw()
	NewShape("square").Draw()
}

func abstractFactoryDemo() {
	fmt.Println("Abstract Factory")
	var f UIFactory = WindowsFactory{}
	f.CreateButton().Click()
	f.CreateCheckbox().Check()
}

func builderDemo() {
	fmt.Println("Builder")
	h := new(HouseBuilder).
		Walls(4).
		Doors(2).
		Garden(true).
		Build()
	fmt.Printf("%+v\n", h)
}

func prototypeDemo() {
	fmt.Println("Prototype")
	orig := Document{"Go Notes", 120}
	copy := orig.Clone()
	copy.Title = "Go Notes Copy"
	fmt.Println(orig)
	fmt.Println(copy)
}

func main() {
	fmt.Println("=== Go Creational Design Patterns ===")

	singletonDemo()
	fmt.Println()

	factoryDemo()
	fmt.Println()

	abstractFactoryDemo()
	fmt.Println()

	builderDemo()
	fmt.Println()

	prototypeDemo()
}
