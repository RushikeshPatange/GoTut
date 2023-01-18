package main

import "fmt"

func main() {
	fmt.Println("Welcome to learn ifElse...")

	age := 18

	if age < 18 {
		fmt.Println("Not elligible for vote")
	} else if age > 18 {
		fmt.Println("Eligibal for vote")
	} else {
		fmt.Println("Vote is an identity of voter")
	}

	// if err != nil {
	// 	return nil, err
	// }

	if num := 10; num > 8 {
		fmt.Println("Number is greater than ten") // useful in webrequest handling
	}

}
