// The program below moves all zeroes to the end of a given array.

package main

import "fmt"

func moveZerosToEnd(arr []int) {
	left, right := 0, len(arr)-1

	for left < right {
		// Move the left pointer until a zero is found
		for left < len(arr) && arr[left] != 0 {
			left++
		}

		// Move the right pointer until a non-zero is found
		for right > 0 && arr[right] == 0 {
			right--
		}

		// Swap the elements if the pointers haven't crossed
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
}

func main() {
	arr := []int{24, 0, 7, 0, 4, 0, 7, -5, 8, 2}

	fmt.Println("The original array is:", arr)

	moveZerosToEnd(arr)

	fmt.Println("The new array is:", arr)
}