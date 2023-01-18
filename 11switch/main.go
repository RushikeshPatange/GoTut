package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Local().UnixNano())
	fmt.Println("Welcome to switch case learning....")
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Mo. is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can open")
		fallthrough
	case 2:
		fmt.Println("You can move 2 step")
	case 3:
		fmt.Println("You can move 3 step")
	case 4:
		fmt.Println("You can move 4 step")
	case 5:
		fmt.Println("You can move 5 step")
	case 6:
		fmt.Println("You can move 6 step")

	default:
		fmt.Println("I missed number,roll it again")
	}

}
