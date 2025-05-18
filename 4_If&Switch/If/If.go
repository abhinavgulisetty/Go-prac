package main

import "fmt"

func main() {
	//Declare variable and use it in if statement
	if condition := true; condition {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")

	}
	//Normal if statement
	if true {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}
	//Normal if statement with else if
	if false {
		fmt.Println("Condition is false")
	} else if true {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}
}
