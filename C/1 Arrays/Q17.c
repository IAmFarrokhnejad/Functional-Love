//The program below calculates the expected sum of the sequence up to the length of the array, then subtract the sum of the array's elements to find the missing number. 
// This approach efficiently identifies the missing number without requiring additional space or complex algorithms.

#include <stdio.h>

int findMissingNum(int *arr1, int ar_size) { // Function to find the missing number
    int i, sum = 0, n = ar_size + 1;


    for (i = 0; i < ar_size; i++) {     // Sum of all elements in the array
        sum = sum + arr1[i];
    }

    // Calculating the missing number using the sum of the first 'n' natural numbers formula
    return (n * (n + 1)) / 2 - sum;
}

int main() {
    int i;
    int arr1[] = {1, 3, 4, 7, 5, 6, 9, 8};
    int ctr = sizeof(arr1) / sizeof(arr1[0]);

    // Displaying the given array
    printf("The given array is :  ");
    for (i = 0; i < ctr; i++) {
        printf("%d  ", arr1[i]);
    }
    printf("\n");

    // Finding and displaying the missing number
    printf("The missing number is : %d \n", findMissingNum(arr1, ctr));

    return 0;
}