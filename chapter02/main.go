package main

import (
	"log"
	"os"

	_ "sammietocat/2016.go.in.action/chapter02/matchers"
	"sammietocat/2016.go.in.action/chapter02/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
