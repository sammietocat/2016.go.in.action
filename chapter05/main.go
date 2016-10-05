// Listing 5.46
// Sample program to show how you cannot always get the address of a value
package main

import (
	"fmt"
)

// duration is a type with a base type of int.
type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// main is the entry point for the application.
func main() {
	duration(42).pretty()

	/*
		./main.go:18: cannot call pointer method on duration(42)
		./main.go:18: cannot take the address of duration(42)
		exit status 2
	*/
}
