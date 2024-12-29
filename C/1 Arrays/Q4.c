// The program below counts the number of duplicate elements in an array. 
// The program takes a specified number of integer inputs, stores them in an array, and then determines and displays how many elements appear more than once.

#include <stdio.h>

int main()
{
    int arr[400];
    int n, mm = 1, ctr = 0;
    int i, j; 
    
    printf("Enter the number of elements to be stored in the array :");
    scanf("%d", &n);

    printf("Input %d elements in the array :\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i);
        scanf("%d", &arr[i]);
    }

    for (i = 0; i < n; i++) {
        for (j = i + 1; j < n; j++) {
            if (arr[i] == arr[j]) {
                ctr++;
                break;  // Exit the inner loop as soon as a duplicate is found
            }
        }
    }
    
    printf("Total number of duplicates found in the array: %d\n", ctr);
    return 0;
}