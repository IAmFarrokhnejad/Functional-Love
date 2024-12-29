// The program below involves copying the elements from one array to another. 
// The program takes a specified number of integer inputs to store in the first array, then copies these elements to a second array, and finally displays the contents of both arrays.

#include <stdio.h>

int main()
{
    int arr1[400], arr2[400];
    int i, n;

    printf("Enter the number of elements to be stored in the array :");
    scanf("%d", &n);
    
    printf("Input %d elements in the array :\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i +1);
        scanf("%d", &arr1[i]); 
    }

    for (i = 0; i < n; i++) {
        arr2[i] = arr1[i];
    }

    printf("\nThe elements stored in the first array are :\n");
    for (i = 0; i < n; i++) {
        printf("% 5d", arr1[i]);
    }
  
    printf("\n\nThe elements copied into the second array are :\n");
    for (i = 0; i < n; i++) {
        printf("% 5d", arr2[i]);
    }
	return 0;
}