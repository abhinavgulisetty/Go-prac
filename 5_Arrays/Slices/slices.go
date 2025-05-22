package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums1 []int // Declare a slice of integers
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 2) // Append elements to the slice
	fmt.Println(nums1)

	nums2 := make([]int, 3, 4) // Create a slice of integers with length 5 and capacity 10
	nums2[0] = 1
	nums2[1] = 2
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))
	nums2 = append(nums2, 3)
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))
	nums2 = append(nums2, 4)
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2)) // Capacity will be doubled when it exceeds the original capacity

	nums3 := []int{1, 2, 3, 4, 5} // Declare and initialize a slice of integers
	nums3 = append(nums3, 6)      // Append an element to the slice
	fmt.Println(nums3)            // Print the slice

	//Copying a slice
	nums4 := []int{1, 2, 3, 4, 5}
	nums5 := make([]int, len(nums4))
	copy(nums5, nums4)
	fmt.Println(nums5) // Print the copied slice
	// Slicing a slice
	nums6 := nums4[1:4] // Slice from index 1 to 3 (exclusive)
	fmt.Println(nums6)  // Print the sliced slice
	// Slicing a slice with default values
	nums7 := nums4[:3] // Slice from index 0 to 2 (exclusive)
	fmt.Println(nums7) // Print the sliced slice
	// Slicing a slice with default values
	nums8 := nums4[2:] // Slice from index 2 to the end
	fmt.Println(nums8) // Print the sliced slice
	// Slicing a slice with default values
	nums9 := nums4[:]  // Slice from index 0 to the end
	fmt.Println(nums9) // Print the sliced slice
	// Slicing a slice with default values
	nums10 := nums4[1:4:4] // Slice from index 1 to 3 (exclusive) with capacity 4
	fmt.Println(nums10)    // Print the sliced slice

	fmt.Println(slices.Equal(nums4, nums5)) // Check if the slices are equal
	fmt.Println(slices.Equal(nums4, nums6)) // Check if the slices are equal
}
