// Variables in golang
package main

import "fmt"

func main() {
	// 1. VARIABLES

	// Variable declaration with no initialization

	var foo string

	// declaration with initialisation

	var foo_two string = "Go is awesome"

	// Ways for Mutliple declarations of variables

	var foo_three, bar string = "Hello", "Rustlangoes"
	// OR
	var (
		foo_four = "Hello"
		bar_one  = "Rustlangoes"
	)

	// Variable decalartion with no type annotation
	// Go compiler can infer some of the types at compile time
	var foo_five = "What is my type?"
	fmt.Printf("foo_five is of type %T", foo_five)

	// shorthand variable declaration with no var keyword and no type annotation
	foo_six := "Shorthand"
	fmt.Printf(":= is for shorthand declaration, var foo_six is of type %T", foo_six)
	// the following print statement is just to silince the compiler
	// static typed lanaguages do not like a program to have unused variables
	fmt.Println(foo, foo_two, foo_three, foo_four, foo_five, foo_six, bar, bar_one)

	//2. Constants

	// constants use the const keyword - they ar efixed values which can never be
	// reassigned
}
