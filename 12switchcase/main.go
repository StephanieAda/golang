package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switchcase")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, you can open")

	case 2:
		fmt.Println("Dice value is 2, you can move 2 paces")

	case 3:
		fmt.Println("Dice value is 3, you can move 3 paces")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4, you can move 4 paces")

	case 5:
		fmt.Println("Dice value is 5, you can move 5 paces")

	case 6:
		fmt.Println("Dice value is 6, you can move 6 paces or play again")

	default:
		fmt.Println("What was that?!")

	}
}
