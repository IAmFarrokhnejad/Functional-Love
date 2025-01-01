//The program below calculates the expected sum of the sequence up to the length of the array, then subtract the sum of the array's elements to find the missing number.

package main

import "fmt"

func findMissingNum(arr []int, arrSize int) int {
	n := arrSize + 1 // Total elements
	expectedSum := n * (n + 1) / 2
	actualSum := 0


	for _, num := range arr { 	// Sum of all elements in the array
		actualSum += num
	}

	return expectedSum - actualSum 	// The missing number
}

func main() {
	arr := []int{1, 3, 4, 7, 5, 2, 9, 8}
	arrSize := len(arr)

	fmt.Println("The given array is:", arr)

	missingNumber := findMissingNum(arr, arrSize)
	fmt.Printf("The missing number is: %d\n", missingNumber)
}