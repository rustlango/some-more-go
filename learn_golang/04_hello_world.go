// initialize go moduel with go mod init command
// go module is a collection og go packages

// always define the package calle dmain because main is the entry point
// of a golang program
package main

// fmt is part of the standard library so we must import it
import "fmt"

// main function is the entry point for the golang program which encapsulates
// the go program
func main() {
	fmt.Println("Hello, Rustlangoes!")
}

// can run code using the go run command (no executable will be produced)
