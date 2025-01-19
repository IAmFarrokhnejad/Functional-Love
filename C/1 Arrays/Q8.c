//The following program separates odd and even integers from a given array into two separate arrays. 
// The program will take a specified number of integer inputs, store them in an array, and then create and display two new arrays: one containing all even elements and the other containing all odd elements.
// Author: Morteza Farrokhnejad
#include <stdio.h>

int main()
{
    int arr1[10], arr2[10], arr3[10];
    int i, j = 0, k = 0, n;

    printf("Enter the number of elements to be stored in the array :");
    scanf("%d", &n);

    printf("Input %d elements in the array :\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i);
        scanf("%d", &arr1[i]);
    }

    for (i = 0; i < n; i++) {
        if (arr1[i] % 2 == 0)
        {
            arr2[j] = arr1[i];
            j++;
        } else {
            arr3[k] = arr1[i];
            k++;
        }
    }

    printf("\nThe Even elements are : \n");
    for (i = 0; i < j; i++) {
        printf("%d ", arr2[i]);
    }

    printf("\nThe Odd elements are :\n");
for (i = 0; i < k; i++) {
        printf("%d ", arr3[i]);
    }
	return 0;
}