package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to converion learning...")
	fmt.Println("Please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	// numtoString, err := strconv.ParseFloat(input, 64) cause \n error
	numtoString, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to given rating,", numtoString+1)
	}
}
