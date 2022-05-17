package main

import "fmt"

func main() {
	// a simple slice
	// a slice is a reference to an array
	// a slice does not store any data, it just describes a section of an array
	// slices are like references to arrays
	// slices can be created with the built-in make function
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(len(a))
}
