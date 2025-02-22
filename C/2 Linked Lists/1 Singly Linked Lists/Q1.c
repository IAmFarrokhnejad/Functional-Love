// The program below demosntrates implementation of a singly linked list. Following functions are included:

// 1. createNodeList(int n): Creates a new linked list with n nodes, where data for each node is input by the user.
// 2. displayList(): Prints all elements in the linked list with arrow separators (e.g., "1 -> 2 -> 3").
// 3. deleteFirstNode(): Removes the first node from the linked list and updates the head.
// 4. deleteMiddleNode(int position): Finds and removes the middle node using the "slow and fast pointer" technique.
// 5. deleteLastNode(): Traverses to the end of the list and removes the last node.
// 6. reverseAndDisplay(): Reverses the order of nodes in the list and displays the result.
// 7. insertNodeBeginning(int num): Adds a new node with value 'num' at the start of the list.
// 8. insertNodeAtEnd(int num): Adds a new node with value 'num' at the end of the list.
// 9. insertNodeMiddle(int num): Adds a new node with value 'num' in the middle of the list.
// 10. searchElement(int value): Searches for a value in the list and prints its position if found.


// Author: Morteza Farrokhnejad
#include <stdio.h>
#include <stdlib.h>

// Structure for a node in a linked list
struct node {
    int num;                // Data of the node
    struct node *next;      // Address of the next node
} *head;                    // Pointer to the starting node (head)

// Function prototypes
void createNodeList(int n);
void displayList();
void reverseAndDisplay();
void insertNodeBeginning(int num);
void insertNodeAtEnd(int num);
void insertNodeMiddle(int num);
void deleteFirstNode();
void deleteMiddleNode(int position);
void deleteLastNode();
void searchElement(int value);


int main() {
    int n;
    printf("Enter the number of nodes: ");
    scanf("%d", &n);

    createNodeList(n);
    
    printf("\nOriginal Linked List Content:\n");
    displayList();
    
    insertNodeBeginning(7);
    printf("\nAfter inserting 7 at the beginning:\n");
    displayList();
    
    insertNodeAtEnd(99);
    printf("\nAfter inserting 99 at the end:\n");
    displayList();
    
    insertNodeMiddle(50);
    printf("\nAfter inserting 50 in the middle:\n");
    displayList();

    deleteFirstNode();
    printf("\nAfter deleting the first node:\n");
    displayList();

    deleteLastNode();
    printf("\nAfter deleting the last node:\n");
    displayList();

    deleteMiddleNode(2);  // Example: Deleting the node at position 2
    printf("\nAfter deleting the middle node at position 2:\n");
    displayList();
    
    printf("\nReversed Linked List Content:\n");
    reverseAndDisplay();

    int searchValue;
    printf("\nInput the element to be searched : ");
    scanf("%d", &searchValue);
    searchElement(searchValue);

    return 0;
}

void createNodeList(int n) {
    struct node *newNode, *tmp;
    int num, i;

    head = NULL;

    for(i = 1; i <= n; i++) {
        newNode = (struct node *)malloc(sizeof(struct node));
        if(newNode == NULL) {
            printf("Memory allocation failed.\n");
            exit(EXIT_FAILURE);
        }

        printf("Enter data for node %d: ", i);
        scanf("%d", &num);

        newNode->num = num;
        newNode->next = NULL;

        if(head == NULL) {
            head = newNode;
        } else {
            tmp->next = newNode;
        }
        tmp = newNode;
    }
}

void displayList() {
    struct node *tmp;
    if(head == NULL) {
        printf("List is empty.\n");
        return;
    }

    tmp = head;
    printf("Linked List: ");
    while(tmp != NULL) {
        printf("%d", tmp->num);
        tmp = tmp->next;
        if(tmp != NULL)
            printf(" -> ");
    }
    printf("\n");
}

void reverseAndDisplay() {
    struct node *prevNode = NULL, *currentNode = head, *nextNode;

    while(currentNode != NULL) {
        nextNode = currentNode->next;
        currentNode->next = prevNode;
        prevNode = currentNode;
        currentNode = nextNode;
    }

    head = prevNode;
    printf("\nAfter Reversal:\n");
    displayList();
}

void insertNodeBeginning(int num) {
    struct node *newNode = (struct node*)malloc(sizeof(struct node));
    if(newNode == NULL) {
        printf("Failed to allocate memory.\n");
        exit(EXIT_FAILURE);
    }

    newNode->num = num;
    newNode->next = head;
    head = newNode;
}

void insertNodeAtEnd(int num) {
    struct node *newNode = (struct node*)malloc(sizeof(struct node));
    if(newNode == NULL) {
        printf("Failed to allocate memory.\n");
        exit(EXIT_FAILURE);
    }

    newNode->num = num;
    newNode->next = NULL;

    if(head == NULL) {
        head = newNode;
        return;
    }

    struct node *tmp = head;
    while(tmp->next != NULL)
        tmp = tmp->next;

    tmp->next = newNode;
}

void insertNodeMiddle(int num) {
    struct node *newNode = (struct node*)malloc(sizeof(struct node));
    if(newNode == NULL) {
        printf("Failed to allocate memory.\n");
        exit(EXIT_FAILURE);
    }

    struct node *slow = head, *fast = head;
    while(fast != NULL && fast->next != NULL) {
        slow = slow->next;
        fast = fast->next->next;
    }

    newNode->num = num;
    newNode->next = slow->next;
    slow->next = newNode;
}

// Function to delete the first node
void deleteFirstNode() {
    if (head == NULL) {
        printf("List is empty, cannot delete.\n");
        return;
    }

    struct node *tmp = head;
    head = head->next;
    free(tmp);
}

// Function to delete a middle node at a given position
void deleteMiddleNode(int position) {
    if (head == NULL) {
        printf("List is empty, cannot delete.\n");
        return;
    }

    struct node *tmp = head, *prev = NULL;
    
    if (position == 1) {
        deleteFirstNode();
        return;
    }

    int count = 1;
    while (tmp != NULL && count < position) {
        prev = tmp;
        tmp = tmp->next;
        count++;
    }

    if (tmp == NULL) {
        printf("Position out of range.\n");
        return;
    }

    prev->next = tmp->next;
    free(tmp);
}

// Function to delete the last node
void deleteLastNode() {
    if (head == NULL) {
        printf("List is empty, cannot delete.\n");
        return;
    }

    if (head->next == NULL) { // If only one node exists
        free(head);
        head = NULL;
        return;
    }

    struct node *tmp = head, *prev = NULL;
    while (tmp->next != NULL) {
        prev = tmp;
        tmp = tmp->next;
    }

    prev->next = NULL;
    free(tmp);
}

void searchElement(int value) {
    struct node *temp = head;
    int position = 1;
    int found = 0;

    if(head == NULL) {
        printf("List is empty.\n");
        return;
    }

    while(temp != NULL) {
        if(temp->num == value) {
            found = 1;
            printf("Element found at node %d\n", position);
            return;
        }
        temp = temp->next;
        position++;
    }

    if(!found) {
        printf("Element not found in the list.\n");
    }
}