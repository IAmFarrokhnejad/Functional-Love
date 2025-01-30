// Incomplete

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

// createNodeList creates a linked list with user-provided data
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

// displayList prints the contents of the linked list
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

// reverseAndDisplay reverses the linked list and displays its content
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

// insertNodeAtBeginning adds a new node at the beginning of the linked list
func insertNodeAtBeginning(head *Node, num int) *Node {
	newNode := &Node{data: num, next: head}
	return newNode
}

// insertNodeAtEnd adds a new node at the end of the linked list
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

// insertNodeAtMiddle inserts a node in the middle of the linked list
func insertNodeAtMiddle(head *Node, num int) *Node {
	if head == nil || head.next == nil {
		// If the list is empty or has only one node, insert at the beginning
		return insertNodeAtBeginning(head, num)
	}

	slow, fast := head, head
	var prev *Node

	// Use slow and fast pointers to find the middle
	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	// Insert the new node after prev (or before slow)
	newNode := &Node{data: num, next: slow}
	if prev != nil {
		prev.next = newNode
	} else {
		head = newNode
	}

	return head
}

// deleteFirstNode removes the first node from the linked list
func deleteFirstNode(head *Node) *Node {
    if head == nil {
        fmt.Println("List is empty, nothing to delete.")
        return nil
    }
    
    newHead := head.next
    head.next = nil // Clean up the reference
    return newHead
}

// deleteLastNode removes the last node from the linked list
func deleteLastNode(head *Node) *Node {
    if head == nil {
        fmt.Println("List is empty, nothing to delete.")
        return nil
    }
    
    // If there's only one node
    if head.next == nil {
        return nil
    }
    
    current := head
    // Traverse to the second-to-last node
    for current.next.next != nil {
        current = current.next
    }
    
    // Remove the last node
    current.next = nil
    return head
}

// deleteMiddleNode removes a node from the middle of the linked list
func deleteMiddleNode(head *Node) *Node {
    if head == nil || head.next == nil {
        fmt.Println("List is too short to delete middle node.")
        return head
    }
    
    // If there are only two nodes, remove the second one
    if head.next.next == nil {
        head.next = nil
        return head
    }
    
    // Use slow and fast pointers to find the middle
    slow, fast := head, head
    var prev *Node
    
    // Find the middle node
    for fast != nil && fast.next != nil {
        fast = fast.next.next
        prev = slow
        slow = slow.next
    }
    
    // Delete the middle node by skipping it
    prev.next = slow.next
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

	// Insert node at the beginning
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

	// Insert node at the end
	fmt.Print("\nEnter data for the new node to insert at the end: ")
	_, err = fmt.Scan(&newNum)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}
	head = insertNodeAtEnd(head, newNum)

	fmt.Println("\nLinked List after adding new node at the end:")
	displayList(head)

	// Insert node in the middle
	fmt.Print("\nEnter data for the new node to insert in the middle: ")
	_, err = fmt.Scan(&newNum)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}
	head = insertNodeAtMiddle(head, newNum)

	fmt.Println("\nLinked List after adding new node in the middle:")
	displayList(head)

	// Reverse and display the list
	head = reverseAndDisplay(head)
	// CALL THE NEWLY ADDED FUNCTIONS HERE LATER 
}