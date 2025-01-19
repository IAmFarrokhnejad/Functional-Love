// The program below finds the largest, second largest, and the third largest elements of an array without sorting the said array.
// Author: Morteza Farrokhnejad

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input the size of the array:")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 3 {
		fmt.Println("Please enter a valid size (at least 3).")
		return
	}

	arr := make([]int, n)
	fmt.Printf("Input %d elements in the array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		arr[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			return
		}
	}

	// Initialize the largest, second largest, and third largest
	largest, secondLargest, thirdLargest := arr[0], -1<<31, -1<<31

	for _, num := range arr {
		if num > largest {
			thirdLargest = secondLargest
			secondLargest = largest
			largest = num
		} else if num > secondLargest && num < largest {
			thirdLargest = secondLargest
			secondLargest = num
		} else if num > thirdLargest && num < secondLargest {
			thirdLargest = num
		}
	}

	if secondLargest == -1<<31 || thirdLargest == -1<<31 {
		fmt.Println("Array does not have enough unique elements for second and third largest.")
		return
	}

	fmt.Printf("\nLargest: %d\n", largest)
	fmt.Printf("Second Largest: %d\n", secondLargest)
	fmt.Printf("Third Largest: %d\n", thirdLargest)
}