// Listing 5.71
// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"sammietocat/2016.go.in.action/chapter05/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type User from the entities package.
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	/*
	  ./main.go:17: unknown entities.User field 'email' in struct literal
	   exit status 2
	*/

	fmt.Printf("User: %v\n", u)
}
