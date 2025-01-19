// The program below creates and displays a singly linked list. "reverseAndDisplay" function reverses the list and displays it. "insertNodeAtBeginning" function inserts a new node at the beginning of the list.
// "insertNodeAtBeginning" function inserts a new node at the end of the linked list.
// Author: Morteza Farrokhnejad
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node represents a single node in the linked list
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

func insertNodeAtBeginning(head *Node, num int) *Node {
	newNode := &Node{data: num, next: head}
	return newNode
}

func insertNodeAtEnd(head *Node, num int) *Node {
	newNode := &Node{data: num, next: nil}

	if head == nil {
		return newNode // Handle empty list
	}

	current := head
	for current.next != nil {
		current = current.next // Traverse to the end of the list
	}
	current.next = newNode // Attach the new node at the end

	return head
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

	fmt.Print("\nEnter data for the new node to insert at the beginning: ")
	var newNum int
	_, err = fmt.Scan(&newNum)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}
	head = insertNodeAtBeginning(head, newNum)

	fmt.Println("\nLinked List after adding new node at the beginning:")
	displayList(head)

	fmt.Print("\nEnter data for the new node to insert at the end: ")
	_, err = fmt.Scan(&newNum)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}
	head = insertNodeAtEnd(head, newNum)

	fmt.Println("\nLinked List after adding new node at the end:")
	displayList(head)

	head = reverseAndDisplay(head)
}