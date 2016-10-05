// Listing 5.64
// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"

	"sammietocat/2016.go.in.action/chapter05/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	counter := counters.alertCounter(10)

	/*
		./main.go:15: cannot refer to unexported name counters.alertCounter
		./main.go:15: undefined: counters.alertCounter
		exit status 2
	*/

	fmt.Printf("Counter: %d\n", counter)
}
