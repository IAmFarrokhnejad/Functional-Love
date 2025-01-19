// The program below checks whether an array is a subset of another array.
// Author: Morteza Farrokhnejad

package main

import "fmt"

// Function to check if arr2 is a subset of arr1
func isSubset(arr1, arr2 []int) bool {
	elementMap := make(map[int]bool)

	// Populate the map with elements from arr1
	for _, elem := range arr1 {
		elementMap[elem] = true
	}

	// Check if all elements of arr2 exist in arr1
	for _, elem := range arr2 {
		if !elementMap[elem] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{4, 8, 2, 4, 6, 9, 5, 0, 2}
	arr2 := []int{5, 4, 2, 0, 6, 2}

	fmt.Println("The original array is:", arr1)
	fmt.Println("The second array is:", arr2)

	if isSubset(arr1, arr2) {
		fmt.Println("The second array is a subset of the first array.")
	} else {
		fmt.Println("The second array is not a subset of the first array.")
	}
}