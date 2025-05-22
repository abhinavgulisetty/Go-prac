package main

import "fmt"

func main() {
	//Better for fixed size
	//Constant access time
	var nums [5]int
	nums[0] = 1
	//Printing the array
	fmt.Println(nums)
	//Length of the array
	fmt.Println(len(nums))
	//Array of strings
	var names [3]string = [3]string{"John", "Doe", "Smith"}
	fmt.Println(names)
	//Another way to declare an array
	nums2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums2)
	//Array of arrays
	var nums3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums3)
}
