package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Printf("%T\t%T\n", arr, slice)
}
