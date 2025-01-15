// The program below creates and displays a singly linked list. "reverseAndDisplay" function reverses the list and displays it. "insertNodeBeginning" function inserts a new node at the beginning of the list.
// "insertNodeAtBeginning" function inserts a new node at the end of the linked list.

#include <stdio.h>
#include <stdlib.h>

// Structure for a node in a linked list
struct node {
    int num;                // Data of the node
    struct node *next;      // Address of the next node
} *head;                    // Pointer to the starting node (head)

void createNodeList(int n);       // Function to create the linked list
void displayList();               // Function to display the linked list
void reverseAndDisplay();         // Function to reverse and display the linked list
void insertNodeBeginning(int num); // Function to insert a node at the beginning
void insertNodeAtEnd(int num);     // Function to insert a node at the end

int main() {
    int n;
    printf("Enter the number of nodes: ");
    scanf("%d", &n);

    createNodeList(n); // Creating the linked list with n nodes
    
    printf("\nOriginal Linked List Content:\n");
    displayList();
    
    insertNodeBeginning(7);
    printf("\nAfter inserting 7 at the beginning:\n");
    displayList();
    
    insertNodeAtEnd(99);
    printf("\nAfter inserting 99 at the end:\n");
    displayList();
    
    printf("\nReversed Linked List Content:\n");
    reverseAndDisplay(); // Reverses and displays the linked list

    return 0;
}

void createNodeList(int n) {
    struct node *newNode, *tmp;
    int num, i;

    head = NULL; // Initialize head as NULL

    for(i = 1; i <= n; i++) {
        newNode = (struct node *)malloc(sizeof(struct node));
        if(newNode == NULL) {
            printf("Memory allocation failed.\n");
            exit(EXIT_FAILURE);
        }

        printf("Enter data for node %d: ", i);
        scanf("%d", &num);

        newNode->num = num;      // Set the data for newNode
        newNode->next = NULL;    // Set the next pointer to NULL

        if(head == NULL) {
            head = newNode;      // First node becomes the head
        } else {
            tmp->next = newNode; // Link the previous node to the new node
        }
        tmp = newNode;           // Move tmp to the new node
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
    struct node *prevNode, *currentNode, *nextNode;

    if(head == NULL) {
        printf("List is empty.\n");
        return;
    }

    prevNode = NULL;
    currentNode = head;
    nextNode = NULL;

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
    struct node *newNode;
    newNode = (struct node*)malloc(sizeof(struct node));
    if(newNode == NULL) {
        printf("Failed to allocate memory.\n");
        exit(EXIT_FAILURE);
    }

    newNode->num = num; // Link the data part
    newNode->next = head; // Link the address part
    head = newNode; // Make newNode the first node
}

void insertNodeAtEnd(int num) {
    struct node *newNode, *tmp;
    newNode = (struct node*)malloc(sizeof(struct node));
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

    tmp = head;
    while(tmp->next != NULL)
        tmp = tmp->next;

    tmp->next = newNode;
}