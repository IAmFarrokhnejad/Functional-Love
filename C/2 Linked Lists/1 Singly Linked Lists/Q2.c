// The program below demosntrates implementation of a singly linked list. Following functions are included:

// 1. newNode(int data): Creates a new node with given data.
// 2. detect_and_remove_Loop(struct Node* head): Detects and removes loop in a linked list.
// 3. displayList(struct Node* head): Prints a linked list.
// 4. displayList_char(struct Node* head): Prints a linked list of characters.
// 5. isPalindrome(struct Node* head): Checks if the linked list is palindrome.

// Author: Morteza Farrokhnejad

#include<stdio.h>
#include <stdlib.h>
#include <stdbool.h>

struct Node {
    int data;
    struct Node* next;
};


struct Node* newNode(int data) {
    struct Node* node = (struct Node*) malloc(sizeof(struct Node));
    node->data = data;
    node->next = NULL;
    return node;
}

// Function to detect and remove loop in a linked list
void detect_and_remove_Loop(struct Node* head) {
    struct Node* slow = head;
    struct Node* fast = head;

    // Use Floyd's cycle-finding algorithm to detect a loop
    while (fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast) break; // If a loop is detected, break the loop
    }

    // If a loop is detected
    if (slow == fast) {
        slow = head;
        while (slow->next != fast->next) {
            slow = slow->next;
            fast = fast->next;
        }

        // Break the loop by setting the next pointer of the loop node to NULL
        fast->next = NULL;
    }
}

void displayList(struct Node* head) {
    while (head) {
        printf("%d ", head->data);
        head = head->next;
    }
    printf("\n");
}

// Function to print a linked list of characters
void displayList_char(struct Node* head) {
    while (head) {
        printf("%c", head->data);
        head = head->next;
    }
    printf("\n");
}

// Function to reverse a linked list
struct Node* reverseList(struct Node* head) {
    struct Node* current = head;
    struct Node* next = NULL;
    struct Node* prev = NULL;

    while (current) {
        next = current->next;
        current->next = prev;
        prev = current;
        current = next;
    }

    return prev;
}

bool isPalindrome(struct Node* head) {
    struct Node* slow = head;
    struct Node* fast = head;
    struct Node* prev_slow = head;
    struct Node* midnode = NULL;
    bool res = true;

    // Find the middle node
    while (fast && fast->next) {
        fast = fast->next->next;
        prev_slow = slow;
        slow = slow->next;
    }

    // If the number of nodes is odd, skip the middle node
    if (fast != NULL) {
        midnode = slow;
        slow = slow->next;
    }

    // Reverse the second half of the linked list
    slow = reverseList(slow);
    fast = head;

    // Compare the first and second half of the linked list
    while (slow != NULL) {
        if (fast->data != slow->data) {
            res = false;
            break;
        }

        fast = fast->next;
        slow = slow->next;
    }

    // Reverse the second half of the linked list again
    slow = reverseList(slow);

    // If the number of nodes is odd, set the next pointer of the middle node to NULL
    if (midnode != NULL) {
        prev_slow->next = midnode;
        midnode->next = slow;
    } else
        prev_slow->next = slow;

    return res;
}

int main() {
    // Creating a singly linked list
    struct Node* head = newNode(1);
    head->next = newNode(2);
    head->next->next = newNode(3);
    head->next->next->next = newNode(4);
    head->next->next->next->next = newNode(5);
    head->next->next->next->next->next = newNode(1);

    printf("Original singly linked list:\n");
    displayList(head);

    // Creating a loop in the linked list
    printf("\nCreate the loop:\n");
	head->next->next->next->next->next = head->next->next;

    // Removing the loop from the linked list
	printf("\n\nAfter removing the said loop:\n");
    detect_and_remove_Loop(head);
    displayList(head);
    if (isPalindrome(head))
        printf("Linked list is a palindrome.\n");
    else
        printf("Linked list is not a palindrome.\n");

    // Example 1 for checking if palindrome: Integer palindrome linked list
	struct Node* head2 = newNode(1);
    head2->next = newNode(2);
    head2->next->next = newNode(2);
    head2->next->next->next = newNode(1);
    printf("\nOriginal Singly List:\n");
	displayList(head2);
    if (isPalindrome(head2))
        printf("Linked list is a palindrome.\n");
    else
        printf("Linked list is not a palindrome.\n");

    // Example 2 for checking if palindrome: Character palindrome linked list
	struct Node* head3 = newNode('M');
	head3->next = newNode('A');
    head3->next->next = newNode('D');
    head3->next->next->next = newNode('A');
    head3->next->next->next->next = newNode('M');
    printf("\nOriginal Singly List:\n");
	displayList_char(head3);
    if (isPalindrome(head3))
        printf("Linked list is a palindrome.\n");
    else
        printf("Linked list is not a palindrome.\n");       


    return 0;
}