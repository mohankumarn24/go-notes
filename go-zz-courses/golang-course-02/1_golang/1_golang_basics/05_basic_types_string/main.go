package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Emporio"
	lastName := "Armani"

	// String concatenation
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	// Common string functions
	fmt.Println("ToUpper():", strings.ToUpper(fullName))
	fmt.Println("ToLower():", strings.ToLower(fullName))
	fmt.Println("Contains():", strings.Contains(fullName, "Emporio"))
	fmt.Println("HasPrefix():", strings.HasPrefix(fullName, "Emp"))
	fmt.Println("HasSuffix():", strings.HasSuffix(fullName, "ni"))
	fmt.Println("ReplaceAll():", strings.ReplaceAll(fullName, "Emporio", "Giorgio"))
	fmt.Println("Split():", strings.Split(fullName, " "))
	fmt.Println("TrimSpace():", strings.TrimSpace("  Hello  "))
	fmt.Println("Repeat():", strings.Repeat("*", 10))
}

/*
Full Name: Emporio Armani
ToUpper(): EMPORIO ARMANI
ToLower(): emporio armani
Contains(): true
HasPrefix(): true
HasSuffix(): true
ReplaceAll(): Giorgio Armani
Split(): [Emporio Armani]
TrimSpace(): Hello
Repeat(): **********
*/
