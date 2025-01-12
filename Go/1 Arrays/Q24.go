// The program below returns the counting sort on an array.

package main

import (
	"fmt"
)

// Function to perform counting sort
func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum and minimum values in the array
	max, min := arr[0], arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	rangeSize := max - min + 1 	// Initialize the count array
	count := make([]int, rangeSize)

	// Count the occurrences of each element
	for _, val := range arr {
		count[val-min]++
	}

	sortedIndex := 0 	// Reconstruct the sorted array
	for i, cnt := range count {
		for cnt > 0 {
			arr[sortedIndex] = i + min
			sortedIndex++
			cnt--
		}
	}

	return arr
}

func main() {
	arr := []int{4, 2, 2, 8, 4, 1, 2, 0, 3, 3, 1}

	fmt.Println("Original array:", arr)

	sortedArr := countingSort(arr)

	fmt.Println("Sorted array:", sortedArr)
}