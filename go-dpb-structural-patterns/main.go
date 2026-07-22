package main

import "fmt"

// ===============================
// Go Structural Design Patterns
// Single runnable program
// ===============================

// ---------- Adapter ----------
type Printer interface{ Print() }

type LegacyPrinter struct{}

func (LegacyPrinter) PrintOld() { fmt.Println("Legacy printer") }

type PrinterAdapter struct{ legacy LegacyPrinter }

func (a PrinterAdapter) Print() { a.legacy.PrintOld() }

// ---------- Bridge ----------
type Device interface {
	On()
	Off()
}

type TV struct{}

func (TV) On()  { fmt.Println("TV ON") }
func (TV) Off() { fmt.Println("TV OFF") }

type Remote struct{ device Device }

func (r Remote) PowerOn()  { r.device.On() }
func (r Remote) PowerOff() { r.device.Off() }

// ---------- Composite ----------
type Component interface{ Show() }

type File struct{ name string }

func (f File) Show() { fmt.Println("File:", f.name) }

type Folder struct {
	name     string
	children []Component
}

func (f Folder) Show() {
	fmt.Println("Folder:", f.name)
	for _, c := range f.children {
		c.Show()
	}
}

// ---------- Decorator ----------
type Coffee interface{ Cost() int }

type BasicCoffee struct{}

func (BasicCoffee) Cost() int { return 5 }

type MilkDecorator struct{ coffee Coffee }

func (m MilkDecorator) Cost() int { return m.coffee.Cost() + 2 }

// ---------- Facade ----------
type CPU struct{}

func (CPU) Start() { fmt.Println("CPU started") }

type Memory struct{}

func (Memory) Load() { fmt.Println("Memory loaded") }

type Computer struct {
	cpu CPU
	mem Memory
}

func (c Computer) Start() {
	c.mem.Load()
	c.cpu.Start()
}

// ---------- Flyweight ----------
type Circle struct{ color string }

var cache = map[string]*Circle{}

func GetCircle(color string) *Circle {
	if c, ok := cache[color]; ok {
		return c
	}
	c := &Circle{color: color}
	cache[color] = c
	return c
}

// ---------- Proxy ----------
type Image interface{ Display() }

type RealImage struct{}

func (RealImage) Display() { fmt.Println("Displaying image") }

type ImageProxy struct{ real *RealImage }

func (p *ImageProxy) Display() {
	if p.real == nil {
		fmt.Println("Lazy loading...")
		p.real = &RealImage{}
	}
	p.real.Display()
}

// ---------- Demo Functions ----------
func adapterDemo() {
	fmt.Println("Adapter")
	var p Printer = PrinterAdapter{legacy: LegacyPrinter{}}
	p.Print()
}

func bridgeDemo() {
	fmt.Println("Bridge")
	r := Remote{device: TV{}}
	r.PowerOn()
	r.PowerOff()
}

func compositeDemo() {
	fmt.Println("Composite")
	root := Folder{
		name: "Documents",
		children: []Component{
			File{"resume.pdf"},
			File{"notes.txt"},
		},
	}
	root.Show()
}

func decoratorDemo() {
	fmt.Println("Decorator")
	var c Coffee = BasicCoffee{}
	c = MilkDecorator{coffee: c}
	fmt.Println("Cost:", c.Cost())
}

func facadeDemo() {
	fmt.Println("Facade")
	pc := Computer{}
	pc.Start()
}

func flyweightDemo() {
	fmt.Println("Flyweight")
	c1 := GetCircle("Red")
	c2 := GetCircle("Red")
	fmt.Println("Shared instance?", c1 == c2)
}

func proxyDemo() {
	fmt.Println("Proxy")
	p := &ImageProxy{}
	p.Display()
	p.Display()
}

func main() {
	fmt.Println("=== Go Structural Design Patterns ===")

	adapterDemo()
	fmt.Println()

	bridgeDemo()
	fmt.Println()

	compositeDemo()
	fmt.Println()

	decoratorDemo()
	fmt.Println()

	facadeDemo()
	fmt.Println()

	flyweightDemo()
	fmt.Println()

	proxyDemo()
}
