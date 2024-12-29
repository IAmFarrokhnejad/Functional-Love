// The program below counts and displays the frequency of each element in an array. 
// The program accepts a specified number of integer inputs, stores them in an array, and then determines and prints how many times each element appears in the array.

#include <stdio.h>

int main()
{
    int arr1[400], fr1[400];
    int n, i, j, ctr;

    printf("Enter the number of elements to be stored in the array :");
    scanf("%d", &n);

    printf("Enter %d elements in the array :\n", n);
    for (i = 0; i < n; i++) {
        printf("element - %d : ", i + 1);
        scanf("%d", &arr1[i]);
        fr1[i] = -1; // Initialize frequency array with -1
    }

    for (i = 0; i < n; i++) {
        ctr = 1; // Counter for each element
        for (j = i + 1; j < n; j++) {
            if (arr1[i] == arr1[j]) {
                ctr++;     // Increment counter for matching elements
                fr1[j] = 0; // Mark the duplicate element's frequency as 0
            }
        }

        // If frequency array value is not marked as 0, set it to the counter
        if (fr1[i] != 0)
            fr1[i] = ctr;
    }

    printf("\nThe frequency of all elements of the array : \n");
    for (i = 0; i < n; i++)
    {
        if (fr1[i] != 0)
            printf("%d occurs %d times\n", arr1[i], fr1[i]);

	}
	return 0;
}