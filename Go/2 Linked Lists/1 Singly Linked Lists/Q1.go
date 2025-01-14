// The program below creates and displays a singly linked list. "reverseAndDisplay" function reverses the list and displays it.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node represents a single node in the list
type Node struct {
	data int
	next *Node
}

func createNodeList(n int) *Node {
	if n <= 0 {
		fmt.Println("Number of nodes must be greater than 0.")
		return nil
	}

	var head, tail *Node
	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i <= n; i++ {
		fmt.Printf("Enter data for node %d: ", i)
		scanner.Scan()
		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			i-- // retry the current iteration
			continue
		}

		newNode := &Node{data: num, next: nil}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
	return head
}

func displayList(head *Node) {
	if head == nil {
		fmt.Println("List is empty.")
		return
	}

	fmt.Println("Linked List Content:")
	for current := head; current != nil; current = current.next {
		fmt.Printf("Data = %d\n", current.data)
	}
}

func reverseAndDisplay(head *Node) *Node {
	var prev, next *Node
	current := head

	for current != nil {
		next = current.next // Save the next node
		current.next = prev // Reverse the link
		prev = current      // Move prev to the current node
		current = next      // Move to the next node
	}

	fmt.Println("\nReversed Linked List Content:")
	displayList(prev)

	return prev
}

func main() {
	fmt.Print("Enter the number of nodes: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid positive integer.")
		return
	}

	head := createNodeList(n)
	fmt.Println("\nOriginal Linked List Content:")
	displayList(head)

	head = reverseAndDisplay(head)
}