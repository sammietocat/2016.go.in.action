// Listing 5.74
// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"sammietocat/2016.go.in.action/chapter05/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported
	// inner type.
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
