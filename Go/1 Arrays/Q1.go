// This program reads a user-defined number of integer values into an array and then display these values in reverse order.
// After storing the values, the program first prints them in the original order and then prints them in the reversed order.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the number of elements to store in the array:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	array := make([]int, n)
	fmt.Printf("Input %d elements in the array:\n", n)

	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		array[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("\nThe values stored in the array are:")
	for _, value := range array {
		fmt.Printf("%5d", value)
	}
	fmt.Println()

	fmt.Println("\nThe values stored in the array in reverse are:")
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%5d", array[i])
	}
}