package main

import "fmt"

//For multiple constants
const (
	PI       = 3.14
	Language = "Go"
)

//f := "Golang"  This is a variable, not a constant
//This will cause an error because f is not a constant

var test string = "test" //This is allowed, but not recommended
// test = "test2"  This will cause an error because test is a variable

func main() {
	const a = 10
	const b = "Go"
	// b = "Golang"  This will cause an error because b is a constant
	fmt.Println(a, b, PI, Language)
}
