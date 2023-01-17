package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slicers learning....")
	fruitList := []string{"Mango", "Apple", "Banana"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("fruitList is ", fruitList)

	fruitList = append(fruitList, "orange", "strawberry")
	fmt.Println("fruitList appended is ", fruitList)

	fruitList = append(fruitList[1:3]) // [firstIndex,lastIndex{Exclusive})
	fmt.Println("fruitList from 1 to 2 is ", fruitList)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	highScores := make([]int, 4) // static mem allocation
	highScores[1] = 111
	highScores[0] = 222
	highScores[2] = 333
	highScores[3] = 444
	// highScores[4] = 555 // cause error: anic: runtime error: index out of range [4] with length 4

	fmt.Println("Highescore slice is ", highScores)

	highScores = append(highScores, 555, 666, 777) // reallocate an entire slice

	fmt.Println("Highescore slice after append function is ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("Sorted highScores slice is ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//  how to remove value from slices based on index

	courses := []string{"reactjs", "nodejs", "js", "py", "cpp"}
	index := 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("Coursenlist after removing js from slice is ", courses)

}
