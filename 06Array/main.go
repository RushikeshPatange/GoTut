package main

import "fmt"

func main() {

	fmt.Println("Welcome to array learning....")

	var sportsIlove [4]string // we have to always specify size of array

	sportsIlove[0] = "cricket"
	sportsIlove[1] = "badmenton"
	sportsIlove[2] = "Swiming"

	fmt.Println("Sport list is", sportsIlove)
	fmt.Println("Sport list is", len(sportsIlove))

	marathiFestival := [5]string{"ganapati", "Diwali", "gudipadva"}
	fmt.Println("Our marathi festival list is", marathiFestival)

}
