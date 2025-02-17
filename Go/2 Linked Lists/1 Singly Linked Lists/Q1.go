// The program below demosntrates implementation of a singly linked list. Following functions are included:

// 1. createNodeList(n int): Creates a new linked list with n nodes, where data for each node is input by the user.
// 2. displayList(head *Node): Prints all elements in the linked list with arrow separators (e.g., "1 -> 2 -> 3").
// 3. deleteFirstNode(head *Node): Removes the first node from the linked list and updates the head.
// 4. deleteMiddleNode(head *Node): Finds and removes the middle node using the "slow and fast pointer" technique.
// 5. deleteLastNode(head *Node): Traverses to the end of the list and removes the last node.
// 6. reverseAndDisplay(head *Node): Reverses the order of nodes in the list and displays the result.
// 7. insertNodeAtBeginning(head *Node, num int): Adds a new node with value 'num' at the start of the list.
// 8. insertNodeAtEnd(head *Node, num int): Adds a new node with value 'num' at the end of the list.
// 9. insertNodeAtMiddle(head *Node, num int): Adds a new node with value 'num' in the middle of the list.
// 10. searchElement(head *Node, value int): Searches for a value in the list and prints its position if found.

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

	fmt.Print("Linked List Content: ")
	current := head
	for current != nil {
		fmt.Printf("%d", current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println()
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
	fmt.Println("First node deleted successfully.")
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
		fmt.Println("Last node (only node) deleted successfully.")
		return nil
	}

	current := head
	// Traverse to the second-to-last node
	for current.next.next != nil {
		current = current.next
	}

	// Remove the last node
	current.next = nil
	fmt.Println("Last node deleted successfully.")
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
		fmt.Println("Middle node deleted successfully.")
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
	fmt.Println("Middle node deleted successfully.")
	return head
}

// searchElement searches for a value in the linked list and prints its position
func searchElement(head *Node, value int) {
	if head == nil {
		fmt.Println("List is empty.")
		return
	}

	current := head
	position := 1
	found := false

	for current != nil {
		if current.data == value {
			fmt.Printf("Element found at node %d\n", position)
			found = true
			break
		}
		current = current.next
		position++
	}

	if !found {
		fmt.Println("Element not found in the list.")
	}
}

func main() {
	fmt.Print("Enter the number of nodes: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid positive integer.")
		return
	}
	
	// Clear the input buffer
	bufio.NewReader(os.Stdin).ReadString('\n')

	head := createNodeList(n)
	fmt.Println("\nOriginal Linked List Content:")
	displayList(head)

	// Test insert operations
	head = insertNodeAtBeginning(head, 7)
	fmt.Println("\nAfter inserting 7 at the beginning:")
	displayList(head)

	head = insertNodeAtEnd(head, 99)
	fmt.Println("\nAfter inserting 99 at the end:")
	displayList(head)

	head = insertNodeAtMiddle(head, 50)
	fmt.Println("\nAfter inserting 50 in the middle:")
	displayList(head)

	// Test delete operations
	head = deleteFirstNode(head)
	fmt.Println("\nAfter deleting first node:")
	displayList(head)

	head = deleteMiddleNode(head)
	fmt.Println("\nAfter deleting middle node:")
	displayList(head)

	head = deleteLastNode(head)
	fmt.Println("\nAfter deleting last node:")
	displayList(head)

	// Test search operation
	fmt.Print("\nEnter the element to be searched: ")
	var searchValue int
	_, err = fmt.Scan(&searchValue)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}
	searchElement(head, searchValue)

	// Finally, reverse the list
	head = reverseAndDisplay(head)
}