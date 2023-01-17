package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointer learning...")

	// var ptr *int

	// fmt.Println("Pointer value ", ptr) // <nil>

	number := 25

	// One way to initilize it
	// var ptr *int
	// ptr = &number

	// other way to initilize it

	var ptr = &number

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value inside pointer is ", *ptr)

	*ptr *= 2

	fmt.Println("New value of variable numbetr is ", number)
	return

}
