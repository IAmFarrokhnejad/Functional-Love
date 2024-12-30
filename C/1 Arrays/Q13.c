// The program below finds the largest, second largest, and the third largest elements of an array without sorting the said array.


#include <stdio.h>
#include <limits.h> // For INT_MIN

int main() {
    int arr[100];
    int n, i;
    int largest, second_largest, third_largest;

    printf("Enter the number of elements in the array: ");
    scanf("%d", &n);

    if (n < 3) {
        printf("The array must have at least 3 elements.\n");
        return 1; // Exit the program
    }

    printf("Enter %d elements:\n", n);
    for (i = 0; i < n; i++) {
        printf("Element - %d: ", i + 1);
        scanf("%d", &arr[i]);
    }

    largest = second_largest = third_largest = INT_MIN;
    for (i = 0; i < n; i++) {
        if (arr[i] > largest) {
            third_largest = second_largest;
            second_largest = largest;
            largest = arr[i];
        } else if (arr[i] > second_largest && arr[i] < largest) {
            third_largest = second_largest;
            second_largest = arr[i];
        } else if (arr[i] > third_largest && arr[i] < second_largest) {
            third_largest = arr[i];
        }
    }
   
    if (second_largest == INT_MIN || third_largest == INT_MIN) {  // Check if second and third largest elements were found
        printf("The array doesn't have distinct second and third largest elements.\n");
    } else {
        printf("The largest element is: %d\n", largest);
        printf("The second largest element is: %d\n", second_largest);
        printf("The third largest element is: %d\n", third_largest);
    }
    return 0;
}