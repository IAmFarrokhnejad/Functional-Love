// The program below insert values in the array (unsorted list). It promots the user to enter an integer value and a value to be inserted and its position.
// The program will then display the array after inserting the new value at the given position.
// Author: Morteza Farrokhnejad
#include <stdio.h>

void main()
{
    int arr1[100], i, n, p, x;


    printf("Input the size of array : ");
    scanf("%d", &n);

    printf("Input %d elements in the array in ascending order:\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i + 1);
        scanf("%d", &arr1[i]);
    }

    printf("Input the value to be inserted : ");
    scanf("%d", &x);
    printf("Input the Position, where the value to be inserted :");
    scanf("%d", &p);

    printf("The current list of the array :\n");
    for (i = 0; i < n; i++)
        printf("% 5d", arr1[i]);

    for (i = n; i >= p; i--)    // Move all data to the right side of the array to make space
        arr1[i] = arr1[i - 1];

    arr1[p - 1] = x;    // Insert the new value at the given position

    printf("\n\nAfter Insert the element the new list is :\n");
    for (i = 0; i <= n; i++)
        printf("% 5d", arr1[i]);
    printf("\n\n");
}