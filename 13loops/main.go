package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, value)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			//break
			rougeValue++
			continue
		}
		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at a goat")
}
