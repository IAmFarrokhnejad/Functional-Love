// The program returns the counting sort on an array.
// Author: Morteza Farrokhnejad

#include <stdio.h>

void countingSort(int arr1[], int n, int max) {
    int count[50] = {0}; // An array to store count of elements
    int i, j;

    // Counting occurrences of each element in the array
    for (i = 0; i < n; ++i)
        count[arr1[i]] = count[arr1[i]] + 1;

    printf("After sorting the elements in the array are:  ");

    // Reconstructing the sorted array using count array
    for (i = 0; i <= max; ++i) {
        for (j = 1; j <= count[i]; ++j) {
            printf("%d ", i); // Printing the element based on its count
        }
    }
}

int main() {
    int max = 0;
    int arr1[] = {1, 14, 8, 0, 2, 4, 2, 1, 0, 17, 9, 0, 22};
    int n = sizeof(arr1) / sizeof(arr1[0]);
    int i;

    printf("The given array is :  ");
    for (i = 0; i < n; i++) {
        if (arr1[i] > max)
            max = arr1[i];
        printf("%d  ", arr1[i]);
    }
    printf("\n");

    countingSort(arr1, n, max); 
    return 0;
}