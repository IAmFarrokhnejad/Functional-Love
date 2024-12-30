// The program below sorts an array in both ascending and descending order using the Quicksort algorithm.
// Quicksort is one of the most efficient sorting algorithms with an average time complexity of 𝑂(𝑛 log 𝑛).

#include <stdio.h>

void quicksort(int arr[], int low, int high, int ascending);
int partition(int arr[], int low, int high, int ascending);
void printArray(int arr[], int size);

void quicksort(int arr[], int low, int high, int ascending) {
    if (low < high) {
        // Partitioning index
        int pi = partition(arr, low, high, ascending);

        // Recursively sort elements before and after partition
        quicksort(arr, low, pi - 1, ascending);
        quicksort(arr, pi + 1, high, ascending);
    }
}

int partition(int arr[], int low, int high, int ascending) {
    int pivot = arr[high]; // Pivot element
    int i = low - 1;       // Index of smaller element

    for (int j = low; j < high; j++) {
        if ((ascending && arr[j] < pivot) || (!ascending && arr[j] > pivot)) {
            i++;
            // Swap arr[i] and arr[j]
            int temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
        }
    }

    // Swap arr[i+1] and arr[high] (or pivot)
    int temp = arr[i + 1];
    arr[i + 1] = arr[high];
    arr[high] = temp;

    return i + 1;
}

void printArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int n;
    printf("Enter the number of elements in the array: ");
    scanf("%d", &n);
    int arr[n];

    printf("Enter the elements of the array:\n");
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }

    int ascending[n];
    for (int i = 0; i < n; i++) {
        ascending[i] = arr[i];
    }
    quicksort(ascending, 0, n - 1, 1);

    int descending[n];
    for (int i = 0; i < n; i++) {
        descending[i] = arr[i];
    }
    quicksort(descending, 0, n - 1, 0);

    printf("\nArray in ascending order:\n");
    printArray(ascending, n);

    printf("\nArray in descending order:\n");
    printArray(descending, n);

    return 0;
}