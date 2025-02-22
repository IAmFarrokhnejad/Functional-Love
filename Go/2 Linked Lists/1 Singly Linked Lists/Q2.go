// The program below demosntrates implementation of a singly linked list. Following functions are included:

// 1. newNode(data int) *Node: Creates a new node with given data.
// 2. detectAndRemoveLoop(head *Node): Detects and removes loop in a linked list.
// 3. displayList(head *Node): Prints a linked list.
// 4. displayListChar(head *Node): Prints a linked list of characters.
// 5. isPalindrome(head *Node): Checks if the linked list is palindrome.

// Author: Morteza Farrokhnejad

package main

import (
	"fmt"
)

// Node structure for a singly linked list
type Node struct {
	data int
	next *Node
}

// Creates a new node with the given data
func newNode(data int) *Node {
	return &Node{data: data, next: nil}
}

// Displays the linked list (assumes no cycles)
func displayList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

// Displays the linked list as characters
func displayListChar(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("%c", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

// Detects and removes a loop in the linked list
func detectAndRemoveLoop(head *Node) {
	if head == nil {
		return
	}

	slow, fast := head, head
	loopExists := false

	// Detect loop using Floyd’s cycle-finding algorithm
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast { // Loop detected
			loopExists = true
			break
		}
	}

	// If a loop is found, remove it
	if loopExists {
		slow = head
		var prev *Node

		for slow != fast {
			prev = fast
			slow = slow.next
			fast = fast.next
		}

		// Break the loop
		prev.next = nil
	}
}

// Reverses the linked list and returns the new head
func reverseList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

// Checks if a linked list is a palindrome
func isPalindrome(head *Node) bool {
	if head == nil {
		return true
	}

	// Find the middle of the linked list
	slow, fast := head, head
	var prevSlow *Node

	for fast != nil && fast.next != nil {
		prevSlow = slow
		slow = slow.next
		fast = fast.next.next
	}

	// Handle odd-length lists (skip middle node)
	midNode := slow
	if fast != nil {
		slow = slow.next
	}

	// Reverse second half
	secondHalf := reverseList(slow)

	// Compare first and reversed second halves
	p1, p2 := head, secondHalf
	result := true

	for p2 != nil {
		if p1.data != p2.data {
			result = false
			break
		}
		p1 = p1.next
		p2 = p2.next
	}

	// Restore the original list structure
	reversedAgain := reverseList(secondHalf)
	if midNode != nil {
		prevSlow.next = midNode
		midNode.next = reversedAgain
	} else {
		prevSlow.next = reversedAgain
	}

	return result
}

// Main function to demonstrate functionality
func main() {
	// ----- Example 1: List with a loop -----
	fmt.Println("Example 1: Detect and Remove Loop")
	node1 := newNode(1)
	node2 := newNode(2)
	node3 := newNode(3)
	node4 := newNode(4)
	node5 := newNode(5)

	// Create linked list: 1 -> 2 -> 3 -> 4 -> 5 -> 1 (loop)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node3 // Creates a cycle

	fmt.Println("Loop created (displaying the list would cause an infinite loop).")
	fmt.Println("Removing loop...")
	detectAndRemoveLoop(node1)

	fmt.Println("After removing loop:")
	displayList(node1)

	if isPalindrome(node1) {
		fmt.Println("Linked list is a palindrome.")
	} else {
		fmt.Println("Linked list is not a palindrome.")
	}

	// ----- Example 2: Integer palindrome list -----
	fmt.Println("\nExample 2: Integer Palindrome List")
	head2 := newNode(1)
	node2_2 := newNode(2)
	node3_2 := newNode(2)
	node4_2 := newNode(1)

	head2.next = node2_2
	node2_2.next = node3_2
	node3_2.next = node4_2

	displayList(head2)
	if isPalindrome(head2) {
		fmt.Println("Linked list is a palindrome.")
	} else {
		fmt.Println("Linked list is not a palindrome.")
	}

	// ----- Example 3: Character palindrome list -----
	fmt.Println("\nExample 3: Character Palindrome List")
	head3 := newNode(int('M'))
	node2_3 := newNode(int('A'))
	node3_3 := newNode(int('D'))
	node4_3 := newNode(int('A'))
	node5_3 := newNode(int('M'))

	head3.next = node2_3
	node2_3.next = node3_3
	node3_3.next = node4_3
	node4_3.next = node5_3

	displayListChar(head3)
	if isPalindrome(head3) {
		fmt.Println("Linked list is a palindrome.")
	} else {
		fmt.Println("Linked list is not a palindrome.")
	}
}