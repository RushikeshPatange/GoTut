package main

import "fmt"

// mongoKey := "asdfghjkl"; // can't use walorous operator outside the method
const MongoKey string = "abjfsehk"

func main() {

	fmt.Println("In Variable Section")

	var username0 string = "rushikesh with type declaration"
	fmt.Println(username0)
	fmt.Printf("Variable type :%T\n", username0)

	var username1 = "rushikesh without type declaration" // lexer will put string in front of var name at time of compilation
	fmt.Println(username1)
	fmt.Printf("Variable type :%T\n", username1)

	username2 := "rushikesh with walorous operator"
	fmt.Println(username2)
	fmt.Printf("Variable type :%T\n", username2)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable type :%T\n", isTrue)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable type :%T\n", smallValue)

	var floatVal float64 = 256.1234567890
	fmt.Println(floatVal)
	fmt.Printf("Variable type :%T\n", floatVal)

	fmt.Println(MongoKey)
	fmt.Printf("Variable type :%T\n", MongoKey)

	return
}
