package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Bananas", "Pineapples")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 765
	highScores[3] = 867
	// highScores[4] = 999

	highScores = append(highScores, 555, 474, 345)
	fmt.Println(highScores)

	// sorting out the slices in increasing order
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
