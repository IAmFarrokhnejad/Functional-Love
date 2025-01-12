// The program below checks whether one array is a subset of another array (all elements of the second array are present in the first array).

package main

import "fmt"

func checkIfSubset(arr1, arr2 []int) bool {
	elementMap := make(map[int]bool)

	// Populate the map with elements from arr1
	for _, val := range arr1 {
		elementMap[val] = true
	}

	// Check if all elements of arr2 are in the map
	for _, val := range arr2 {
		if !elementMap[val] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 8, 7, 4, 6, 9, 5, 0, 2}
	arr2 := []int{5, 4, 2, 7, 0, 1, 6, 2}

	fmt.Println("The first array is:", arr1)
	fmt.Println("The second array is:", arr2)

	if checkIfSubset(arr1, arr2) {
		fmt.Println("The second array is a subset of the first array.")
	} else {
		fmt.Println("The second array is not a subset of the first array.")
	}
}
