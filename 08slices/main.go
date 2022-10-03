package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Bananas", "Pineapples")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
}
