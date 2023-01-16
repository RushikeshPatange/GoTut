package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input learning....")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	input, _ := reader.ReadString('\n') // comma, ok || comma, error syntax used as try-catch block emulation
	fmt.Printf("Your name is %s", input)
	fmt.Printf("Your name type is %T", input)
	// fmt.Printf(err.Error())

}
