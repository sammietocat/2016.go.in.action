package main

// Duration   Listing 5.9 Compiler error assigning value of different types
type Duration int64

func main() {
	var dur Duration
	dur = int64(1000)
}
