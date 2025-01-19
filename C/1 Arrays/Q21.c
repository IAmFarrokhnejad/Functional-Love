// The program below checks whether an array is a subset of another array.
// Author: Morteza Farrokhnejad

#include <stdio.h>

int checkSubsetArray(int *arr1 , int arr1_size, int *arr2, int arr2_size) {
    int i, j;
    
    for (i = 0; i < arr2_size; i++) {     // Iterate through elements of arr2[]
        // Search for current arr2[] element in arr1[]
        for (j = 0; j < arr1_size; j++) {
            if (arr2[i] == arr1[j])
                break;
        }
        
        if (j == arr1_size) // If the loop finishes without finding an element, arr2[] is not a subset
            return 0; // Not a subset
    }
    return 1; // arr2[] is a subset of arr1[]
}

int main() {
    int arr1[] = {4, 8, 2, 4, 6, 9, 5, 0, 2};
    int arr2[] = {5, 4, 2, 0, 6, 2};
    int n1 = sizeof(arr1) / sizeof(arr1[0]);
    int n2 = sizeof(arr2) / sizeof(arr2[0]);
    int i;

    printf("The original array is: ");
    for (i = 0; i < n1; i++) {
        printf("%d ", arr1[i]);
    }
    printf("\n");

    printf("The second array is: ");
    for (i = 0; i < n2; i++) {
        printf("%d ", arr2[i]);
    }
    printf("\n");

    if (checkSubsetArray(arr1, n1, arr2, n2))
        printf("The second array is a subset of the first array.\n");
    else
        printf("The second array is not a subset of the first array.\n");
    return 0;
}