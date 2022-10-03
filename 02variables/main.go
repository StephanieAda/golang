package main

import "fmt"

func main() {
	var username bool = false
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.5434357685445323427
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default valies and sime aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type %T \n", anothervariable)

	// implicit type
	var website = "learncood"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T \n", website)

	// no var style
	numberofuser := 300000.89
	fmt.Println(numberofuser)
	fmt.Printf("Variable is of type %T \n", numberofuser)

}
