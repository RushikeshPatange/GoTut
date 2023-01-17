package main

import "fmt"

func main() {

	fmt.Println("Welcome to struct learning...")
	// no inheritance present in Go Lang

	// var Rushikesh = User{79, "Rushikesh Patange", true}
	Rushikesh := User{79, "Rushikesh Patange", true}

	fmt.Printf("User info is %+v\n", Rushikesh)
	fmt.Printf("Username is %v and Status of log in is %v", Rushikesh.Name, Rushikesh.isLoggedIn)
}

type User struct {
	ID         int
	Name       string
	isLoggedIn bool
}
