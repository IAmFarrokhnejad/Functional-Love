// The program below inserts a new value into an already sorted array while maintaining the sorted order.
// The program prompts the user with the number of elements to input, elements in ascending order, and the value to be inserted.
// It then display the array before and after insertion.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MaxSize = 100

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := []int{}

	fmt.Printf("Input number of elements you want to insert (max %d): ", MaxSize)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n > MaxSize {
		fmt.Printf("The size of the array cannot exceed %d. Please try again.\n", MaxSize)
		return
	}

	fmt.Printf("Input %d elements in the array in ascending order:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i)
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, value)
	}

	fmt.Print("Input the value to be inserted: ")
	scanner.Scan()
	inval, _ := strconv.Atoi(scanner.Text())

	fmt.Println("\nThe existing array list is:")
	for _, value := range arr {
		fmt.Printf("%5d", value)
	}

	// Find the position to insert the value
	position := len(arr)
	for i, value := range arr {
		if inval < value {
			position = i
			break
		}
	}

	// Insert the value at the correct position
	arr = append(arr[:position], append([]int{inval}, arr[position:]...)...)

	fmt.Println("\n\nAfter Insert the list is:")
	for _, value := range arr {
		fmt.Printf("%5d", value)
	}
	fmt.Println()
}
// Author: Morteza Farrokhnejad