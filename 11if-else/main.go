package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 38
	var result string

	if loginCount < 11 {
		result = "Regular user"
	} else if loginCount > 11 {
		result = "Stalker"
	} else {
		result = "Exactly safe"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println(("Number is even"))
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 11 {
		fmt.Println("Num is less than 30")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	//}
}
