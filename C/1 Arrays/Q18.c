// The program below finds two array elements whose sum is closest to zero.
// It  iterate through the array, checking pairs of elements, and determine the pair with the smallest absolute sum, then print those elements.
// Author: Morteza Farrokhnejad
#include <stdio.h>
#include <math.h>
#include <stdlib.h>

// Function to find and display the pair of elements whose sum is minimum
void findMinSumPair(int *arr1, int arr_size) {
    int i, j, sum, minSum, min1Pair, min2Pair;

    // Check for valid input array
    if (arr1 == NULL || arr_size < 2)
        return;

    // Initialize variables to track minimum sum and respective pair elements
    min1Pair = arr1[0];
    min2Pair = arr1[1];
    minSum = min1Pair + min2Pair;

    for (i = 0; i < arr_size - 1; i++) { // Loop through the array to find the pair with the minimum sum
        for (j = i + 1; j < arr_size; j++) {
            sum = arr1[i] + arr1[j];
            
            if (abs(sum) < abs(minSum)) { // Update minimum sum and pair elements if the current sum is smaller
                minSum = sum;
                min1Pair = arr1[i];
                min2Pair = arr1[j];
            }
        }
    }
    printf("%d and %d have the closest sum to zero \n", min1Pair, min2Pair);
}

int main() {
    int arr1[] = {2, 7, 63, -80, -2, 19, 84, -77, 4, -1}; 
    int ctr = sizeof(arr1) / sizeof(arr1[0]);
    int i;  

    printf("The original array: ");
    for (i = 0; i < ctr; i++) {
        printf("%d  ", arr1[i]);
    } 
    printf("\n");

    printf("The Pair of elements whose sum is minimum are: \n");
    findMinSumPair(arr1, ctr);
    return 0;
}