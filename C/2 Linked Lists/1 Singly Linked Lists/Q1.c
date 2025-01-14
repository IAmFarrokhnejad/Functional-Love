// The program below creates and displays a singly linked list. "reverseAndDisplay" function reverses the list and displays it. 


#include <stdio.h>
#include <stdlib.h>

// Structure for a node in a linked list
struct node {
    int num;                // Data of the node
    struct node *next;   // Address of the next node
} *head;                // Pointer to the starting node(head)

void createNodeList(int n); // Function to create the linked list
void displayList();         // Function to display the linked list
void reverseAndDisplay();   // Function to reverse and display the linked list

int main() {
    int n;
    printf(" Enter the number of nodes : ");
    scanf("%d", &n);

    createNodeList(n); // Creating the linked list with n nodes
    
    printf("\n Original Linked List Content: \n");
    displayList();
    
    printf("\n Reversed Linked List Content: \n");
    reverseAndDisplay();  // Added the function call here

    return 0;
}

void createNodeList(int n) {
    struct node *newNode, *tmp;
    int num, i;
    head = (struct node *)malloc(sizeof(struct node));     // Allocating memory for the starting node

    if(head == NULL) {
        printf("Memory allocation failed.");
    } else {
        // Reading data for the starting node from user input
        printf(" Enter data for node 1 : ");
        scanf("%d", &num);
        head->num = num;      
        head->next = NULL; // Setting the next pointer to NULL
        tmp = head;

        // Creating n nodes and adding them to the linked list from the second node until the nth
        for(i = 2; i <= n; i++) {
            newNode = (struct node *)malloc(sizeof(struct node));
            if(newNode == NULL) {
                printf("Memory allocation failed.");
                break;
            } else {
                // Reading data for each node from user input
                printf(" Input data for node %d : ", i);
                scanf(" %d", &num);
                newNode->num = num;      // Setting the data for newNode
                newNode->next = NULL; // Setting the next pointer to NULL
                tmp->next = newNode;  // Linking the current node to newNode
                tmp = tmp->next;     // Moving tmp to the next node
            }
        }
    }
}

void displayList() {
    struct node *tmp;
    if(head == NULL) {
        printf(" List is empty.");
    } else {
        tmp = head;
        // Traversing the linked list and printing each node's data
        while(tmp != NULL) {
            printf(" Data = %d\n", tmp->num); // Printing the data of the current node
            tmp = tmp->next;               // Moving to the next node in the list
        }
    }
}

void reverseAndDisplay() {
    struct node *prevNode, *currentNode, *nextNode;
    
    if(head == NULL) {
        printf(" List is empty.");
        return;
    }
    
    // Initialize pointers for reversal
    prevNode = NULL;
    currentNode = head;
    nextNode = NULL;
    
    // Reverse the links between nodes
    while(currentNode != NULL) {
        // Store next node
        nextNode = currentNode->next;
        
        // Reverse the link
        currentNode->next = prevNode;
        
        // Move prevNode and currentNode one step forward
        prevNode = currentNode;
        currentNode = nextNode;
    }
    
    // Update head to point to the last node (new first node)
    head = prevNode;
    
    printf("\n After Reversal:\n");
    displayList();
}