// The program below counts and displays the frequency of each element in an array.
// The program accepts a specified number of integer inputs, stores them in an array, and then determines and prints how many times each element appears in the array.

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

	arr := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	frequency := make(map[int]int)
	for _, element := range arr {
		frequency[element]++
	}

	fmt.Println("\nThe frequency of all elements of the array:")
	for element, count := range frequency {
		fmt.Printf("%d occurs %d times\n", element, count)
	}
}