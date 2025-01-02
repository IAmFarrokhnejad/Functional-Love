// The program below finds two array elements whose sum is closest to zero.

package main

import (
	"fmt"
	"math"
)

func findMinSumPair(arr []int) (int, int, bool) {
	if len(arr) < 2 {
		return 0, 0, false // Not enough elements to find a pair
	}

	min1Pair := arr[0]
	min2Pair := arr[1]
	minSum := min1Pair + min2Pair

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if math.Abs(float64(sum)) < math.Abs(float64(minSum)) {
				minSum = sum
				min1Pair = arr[i]
				min2Pair = arr[j]
			}
		}
	}

	return min1Pair, min2Pair, true
}

func main() {
	arr := []int{10, 1, 24, -10, -2, 19, 22, -24, 4, -1}

	fmt.Println("The original array:", arr)

	// Find and display the pair with a sum closest to zero
	if a, b, found := findMinSumPair(arr); found {
		fmt.Printf("The pair of elements whose sum is minimum are: %d and %d\n", a, b)
	} else {
		fmt.Println("The array does not contain enough elements to find a pair.")
	}
}