// The program below takes a specified number of integer inputs, stores them in an array, and then determines and prints the highest and lowest values among the elements.
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

	fmt.Println("Enter the number of elements:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n == 0 {
		fmt.Println("No elements to process.")
		return
	}

	numbers := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		numbers[i], _ = strconv.Atoi(scanner.Text())
	}

	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Printf("\nThe highest value is: %d\n", max)
	fmt.Printf("The lowest value is: %d\n", min)
}