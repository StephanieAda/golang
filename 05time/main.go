package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 3:04:05 PM Monday"))

	createdDate := time.Date(2020, time.August, 25, 14, 12, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01/02/2006 Monday "))
}
