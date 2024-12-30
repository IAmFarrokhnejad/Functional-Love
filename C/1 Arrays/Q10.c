// The program below inserts a new value into an already sorted array while maintaining the sorted order. 
// The program prompts the user with the number of elements to input, elements in ascending order, and the value to be inserted. 
//It then display the array before and after insertion.
#include <stdio.h>

#define MAX_SIZE 100
int main() {
    int arr1[MAX_SIZE + 1], i, n, p, inval;

    printf("Input number of elements you want to insert (max %d): ", MAX_SIZE);
    scanf("%d", &n);
    // Check if the input size exceeds the maximum allowed size
    if (n > MAX_SIZE) {
        printf("The size of the array cannot exceed %d. Please try again.\n", MAX_SIZE);
        return 1; // Exit the program with an error code
    }

    printf("Input %d elements in the array in ascending order:\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i);
        scanf("%d", &arr1[i]);
    }

    printf("Input the value to be inserted : ");
    scanf("%d", &inval);

    printf("The existing array list is :\n");
    for (i = 0; i < n; i++) {
        printf("% 5d", arr1[i]);
    }
    // Determine the position where the new value will be inserted
    for (i = 0; i < n; i++) {
        if (inval < arr1[i]) {
            p = i;
            break;
        } else {
            p = i + 1;
        }
    }
    // Move all data at the right side of the array to make space
    for (i = n; i >= p; i--) {
        arr1[i + 1] = arr1[i];
    }
    // Insert the new value at the proper position
    arr1[p] = inval;
    // Display the array after insertion
    printf("\n\nAfter Insert the list is :\n");
    for (i = 0; i <= n; i++) {
        printf("% 5d", arr1[i]);
    }
    return 0;
}