package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("Value of Pointer is", ptr)

	myNumber := 23

	var ptrr = &myNumber
	fmt.Println("Vaue of actual pointer is ", ptrr)
	fmt.Println("Vaue of actual pointer is ", *ptrr)

	*ptr = *ptr * 2

	fmt.Println("New value is:", myNumber)

}
