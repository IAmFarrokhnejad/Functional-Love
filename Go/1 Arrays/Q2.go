// The program below reads a specified number of integers into an array and then calculate and print the sum of these elements.
// After storing the input values, the program iterates through the array to compute the total sum and display the result.
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
	fmt.Println("Enter the number of elements to be stored in the array:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	array := make([]int, n)
	sum := 0

	fmt.Printf("Input %d elements in the array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		array[i], _ = strconv.Atoi(scanner.Text())
		sum += array[i]
	}

	fmt.Printf("Sum of all elements stored in the array is: %d\n", sum)
}