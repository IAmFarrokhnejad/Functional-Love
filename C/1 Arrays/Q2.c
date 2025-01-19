// The program below reads a specified number of integers into an array and then calculate and print the sum of these elements. 
// After storing the input values, the program iterates through the array to compute the total sum and display the result.
// Author: Morteza Farrokhnejad
#include <stdio.h>

int main()
{
    int a[400];
    int i, n, sum = 0;

    printf("Enter the number of elements to be stored in the array :");
    scanf("%d", &n);

    printf("Input %d elements in the array :\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i +1);
        scanf("%d", &a[i]);
    }

    for (i = 0; i < n; i++) {
        sum += a[i];
    }

    printf("Sum of all elements stored in the array is : %d\n\n", sum);
	return 0;
}