package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang; no super or parent or class
	princess := User{"Princess", "princess@golang.com", true, 28}
	fmt.Println(princess)
	fmt.Printf("Princess details are: %+v \n", princess)
	fmt.Printf("Princess email is %v and Status is %v\n", princess.Email, princess.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
