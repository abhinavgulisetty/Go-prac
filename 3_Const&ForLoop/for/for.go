package main

import "fmt"

func main() {
	//basic for loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//for loop with initialization, condition and post statement
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}
	//Infinite Loop
	for {
		fmt.Println("Infinite Loop")
		break // to break the infinite loop
	}
}
