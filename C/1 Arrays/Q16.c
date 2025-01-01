// The program below prompts the user to input the size of the matrix and its elements.
// It then calculates the sum of each row and column, and displays the original matrix followed by the sums of its rows and columns as output.

#include <stdio.h>

int main() {
    int i, j, k, arr1[10][10], rsum[10], csum[10], n;


    printf("Input the size of the square matrix : ");
    scanf("%d", &n);

    printf("Input elements in the matrix :\n");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf("element - [%d],[%d] : ", i, j);
            scanf("%d", &arr1[i][j]);
        }
    }

    printf("The original matrix :\n");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf("% 4d", arr1[i][j]);
        }
        printf("\n");
    }

    for (i = 0; i < n; i++) {     // Sum of rows
        rsum[i] = 0;
        for (j = 0; j < n; j++) {
            rsum[i] = rsum[i] + arr1[i][j];
        }
    }

    for (i = 0; i < n; i++) {    // Sum of columns
        csum[i] = 0;
        for (j = 0; j < n; j++) {
            csum[i] = csum[i] + arr1[j][i];
        }
    }

    // Display the matrix with row sums
    printf("The sum of rows and columns of the matrix is :\n");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf("% 4d", arr1[i][j]);
        }
        printf("% 8d", rsum[i]); // Print the sum of each row
        printf("\n");
    }

    // Display the column sums
    printf("\n");
    for (j = 0; j < n; j++) {
        printf("% 4d", csum[j]); // Print the sum of each column
    }
    return 0;
}