package main

import "fmt"

// Listing 5.1 Declatation of a struct type
// user defines a user in the program
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// Listing 5.6 Declating fieldss based on other struct types
// admin represents an admin user with privileges.
type admin struct {
	person user
	level  string
}

// Listing 5.8 Declaration of a new type based on an int64
type Duration int64

func main() {
	// Listing 5.2 Declaration of a variable of the struct type set to its zero value
	// Declare a variable of type user.
	var bill user

	//Listing 5.3/5.4 Declaration of a variable of the struct type using a struct literal
	// Declare a variable of type user and initialize all the fields.
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// Listing 5.5 Creating a struct type value without declaring the field names
	// Declare a variable of type user.
	sammy := user{"sammy", "sammya@email.com", 123, true}

	// Listing 5.7 Using struct literals to create values for fields
	// Declare a variable of type admin
	fred := admin{
		person: user{
			name:       "Fred",
			email:      "fred@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	var duration Duration = 64

	fmt.Println(bill)
	fmt.Println(lisa)
	fmt.Println(sammy)
	fmt.Println(fred)
	fmt.Println(duration)
}
