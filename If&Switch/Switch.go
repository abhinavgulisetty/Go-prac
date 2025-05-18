package main

import "fmt"

func main() {
	// Declare variable and use it in switch statement
	condition := 2
	switch condition {
	case 1:
		fmt.Println("Condition is 1")
	case 2:
		fmt.Println("Condition is 2")
	default:
		fmt.Println("Condition is not 1 or 2")
	}
	// Normal switch statement
	switch {
	case true:
		fmt.Println("Condition is true") // This will execute the first case that is true
	case false:
		fmt.Println("Condition is false")
	default:
		fmt.Println("Condition is not true or false")
	}

	// Normal switch statement with fallthrough
	switch {
	case false:
		fmt.Println("Condition is true")
	case true:
		fmt.Println("Condition is false")
		fallthrough // This will execute the next case as well even if it is false
	default:
		fmt.Println("Condition is not true or false")
	}

	//Type switch statement
	whatAmI := func(i any) { // Use 'any' to accept any type of value or interface{}
		switch i.(type) { // returns the type of the value
		case nil:
			fmt.Println("I am nil") // This will execute if the value is nil
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I am something else")
		}
	}
	whatAmI(42)
	whatAmI("Hello")
	whatAmI(true)
	whatAmI(3.14)
}
