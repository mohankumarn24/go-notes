package main

import "fmt"

// ===================================
// Go Behavioral Design Patterns
// Single runnable program
// ===================================

// ---------- Chain of Responsibility ----------
type Handler interface{ Handle(string) }

type Auth struct{ next Handler }

func (a Auth) Handle(r string) {
	fmt.Println("Auth:", r)
	if a.next != nil {
		a.next.Handle(r)
	}
}

type Logger struct{}

func (Logger) Handle(r string) { fmt.Println("Logger:", r) }

// ---------- Command ----------
type Command interface{ Execute() }
type LightOn struct{}

func (LightOn) Execute() { fmt.Println("Light ON") }

// ---------- Interpreter ----------
type Expr interface{ Interpret() int }
type Number struct{ value int }

func (n Number) Interpret() int { return n.value }

// ---------- Iterator ----------
func iteratorDemo() {
	fmt.Println("Iterator")
	for i, v := range []string{"Go", "Java", "Python"} {
		fmt.Println(i, v)
	}
}

// ---------- Mediator ----------
type ChatRoom struct{}

func (ChatRoom) Send(msg string) { fmt.Println("Chat:", msg) }

// ---------- Memento ----------
type Memento struct{ state string }
type Editor struct{ text string }

func (e Editor) Save() Memento      { return Memento{e.text} }
func (e *Editor) Restore(m Memento) { e.text = m.state }

// ---------- Observer ----------
type Observer interface{ Update(string) }
type Email struct{}

func (Email) Update(msg string) { fmt.Println("Email:", msg) }

type Subject struct{ observers []Observer }

func (s *Subject) Register(o Observer) { s.observers = append(s.observers, o) }
func (s Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

// ---------- State ----------
type State interface{ Handle() }
type Playing struct{}

func (Playing) Handle() { fmt.Println("Playing") }

// ---------- Strategy ----------
type Payment interface{ Pay() }
type Card struct{}

func (Card) Pay() { fmt.Println("Paid by Card") }

// ---------- Template Method ----------
type Worker interface{ Step() }

func Run(w Worker) {
	fmt.Println("Start")
	w.Step()
	fmt.Println("End")
}

type Job struct{}

func (Job) Step() { fmt.Println("Doing work") }

// ---------- Visitor ----------
type Visitor interface{ Visit() }
type ReportVisitor struct{}

func (ReportVisitor) Visit() { fmt.Println("Generating report") }

type File struct{}

func (File) Accept(v Visitor) { v.Visit() }

// ---------- Demo ----------
func chainDemo() {
	fmt.Println("Chain of Responsibility")
	Auth{next: Logger{}}.Handle("Request")
}
func commandDemo() {
	fmt.Println("Command")
	var c Command = LightOn{}
	c.Execute()
}
func interpreterDemo() {
	fmt.Println("Interpreter")
	var e Expr = Number{42}
	fmt.Println(e.Interpret())
}
func mediatorDemo() {
	fmt.Println("Mediator")
	ChatRoom{}.Send("Hello")
}
func mementoDemo() {
	fmt.Println("Memento")
	ed := Editor{text: "Draft"}
	m := ed.Save()
	ed.text = "Changed"
	ed.Restore(m)
	fmt.Println(ed.text)
}
func observerDemo() {
	fmt.Println("Observer")
	s := Subject{}
	s.Register(Email{})
	s.Notify("Order shipped")
}
func stateDemo() {
	fmt.Println("State")
	var s State = Playing{}
	s.Handle()
}
func strategyDemo() {
	fmt.Println("Strategy")
	var p Payment = Card{}
	p.Pay()
}
func templateDemo() {
	fmt.Println("Template Method")
	Run(Job{})
}
func visitorDemo() {
	fmt.Println("Visitor")
	File{}.Accept(ReportVisitor{})
}

func main() {
	fmt.Println("=== Go Behavioral Design Patterns ===")
	chainDemo()
	fmt.Println()
	commandDemo()
	fmt.Println()
	interpreterDemo()
	fmt.Println()
	iteratorDemo()
	fmt.Println()
	mediatorDemo()
	fmt.Println()
	mementoDemo()
	fmt.Println()
	observerDemo()
	fmt.Println()
	stateDemo()
	fmt.Println()
	strategyDemo()
	fmt.Println()
	templateDemo()
	fmt.Println()
	visitorDemo()
}
