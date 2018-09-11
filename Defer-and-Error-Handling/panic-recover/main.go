package main

import (
	"fmt"
)

/**
Here is a quick recap of what we learned in this tutorial,

	What is panic?
	When should panic be used?
	Panic example
	Defer while panicking
	Recover
	Panic, Recover and Goroutines
	Runtime panics
	Getting stack trace after recover
*/

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
