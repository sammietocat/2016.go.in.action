// Listing 5.48
// Sample program to show how polymorphic behavior with interfaces.
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin defines a admin in the program.
type admin struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		a.name,
		a.email)
}

// main is the entry point for the application.
func main() {
	// Create a user value and pass it to  snedNotification.
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// Create a admin value and pass it to  snedNotification.
	lisa := user{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
