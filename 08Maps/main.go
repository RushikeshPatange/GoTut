package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to learn maps...")

	lang := make(map[string]string)

	lang["JS"] = "javascript"
	lang["TS"] = "typescript"
	lang["Py"] = "Python"

	fmt.Println("our lang map is ", lang)

	delete(lang, "TS")

	fmt.Println("After deleting TS key is lang map ", lang)

	//looping over map

	for key, v := range lang {
		fmt.Printf("Value of %s is %s\n", key, v)
	}
}
