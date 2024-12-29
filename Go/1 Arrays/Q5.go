// In this program, we will enter two sorted arrays. While merging them, we will compare the elements of both the arrays and merge them in a sorted manner.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the size of the first array:")
	scanner.Scan()
	n1, _ := strconv.Atoi(scanner.Text())
	a := make([]int, n1)
	fmt.Println("Enter the elements of the first array:")
	for i := 0; i < n1; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Enter the size of the second array:")
	scanner.Scan()
	n2, _ := strconv.Atoi(scanner.Text())
	b := make([]int, n2)
	fmt.Println("Enter the elements of the second array:")
	for i := 0; i < n2; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	c := mergeSortedArrays(a, b)

	fmt.Println("Final array after merging:")
	for _, value := range c {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func mergeSortedArrays(a, b []int) []int {
	n1, n2 := len(a), len(b)
	c := make([]int, 0, n1+n2)

i, j := 0, 0

	for i < n1 && j < n2 {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	c = append(c, a[i:]...)
	c = append(c, b[j:]...)

	return c
}