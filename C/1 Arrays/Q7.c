// The program below takes a specified number of integer inputs, stores them in an array, and then determines and prints the highest and lowest values among the elements.


#include <stdio.h>

int main()
{
    int arr1[400];
    int i, mx, mn, n;

    printf("Enter the number of elements to be stored in the array :");
    scanf("%d", &n);

    printf("Input %d elements in the array :\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i + 1);
        scanf("%d", &arr1[i]);
    }

    mx = arr1[0];
    mn = arr1[0];

    for (i = 1; i < n; i++) {
        if (arr1[i] > mx)
            mx = arr1[i];

        if (arr1[i] < mn)
            mn = arr1[i];
    }

    printf("Maximum element is : %d\n", mx);
    printf("Minimum element is : %d\n\n", mn);
	return 0;
}