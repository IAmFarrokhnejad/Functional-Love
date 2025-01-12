// The program below checks whether one array is a subset of another array (all elements of the second array are present in the first array).
// It does it by iterating through each element of the second array and checking if it exists in the first array. 
// If every element in the second array is found within the first array, then the second array is a subset of the first array.
#include <stdio.h>


int checkIfSubset(int *arr1 , int arr1_size, int *arr2, int arr2_size) {
    int i, j;
    // Iterate through elements of the second array.
    for (i = 0; i < arr2_size; i++) {
        // Search for current arr2 element in arr1.
        for (j = 0; j < arr1_size; j++) {
            // If element found in arr1[], break the loop.
            if (arr2[i] == arr1[j])
                break;
        }
        // If the loop finishes without finding an element, arr2[] is not a subset.
        if (j == arr1_size)
            return 0; // Not a subset
    }
    return 1; // arr2 is a subset of arr1
}

int main() {
    int arr1[] = {4, 2, 7, 1, 6, 9, 5, 0, 2};
    int arr2[] = {5, 4, 2, 0, 1};
    int n1 = sizeof(arr1) / sizeof(arr1[0]);
    int n2 = sizeof(arr2) / sizeof(arr2[0]);
    int i;

    printf("The first array is: ");
    for (i = 0; i < n1; i++) {
        printf("%d ", arr1[i]);
    }
    printf("\n");

    printf("The second array is: ");
    for (i = 0; i < n2; i++) {
        printf("%d ", arr2[i]);
    }
    printf("\n");

    if (checkIfSubset(arr1, n1, arr2, n2))
        printf("The second array is a subset of the first array.\n");
    else
        printf("The second array is not a subset of the first array.\n");

    return 0;
}