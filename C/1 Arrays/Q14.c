// The program below adds two square matrices of the same size. 
//The program prompts the user to input the size of the matrices (less than 5), elements for each matrix, and then calculates the sum of the matrices. It displays the original matrices and their sum as output.

#include <stdio.h>

void main() {
    int arr1[50][50], brr1[50][50], crr1[50][50], i, j, n;

    printf("Enter the size of the square matrix (less than 5): ");
    scanf("%d", &n);

    printf("Enter the elements in the first matrix :\n");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf("element - [%d],[%d] : ", i, j);
            scanf("%d", &arr1[i][j]);
        }
    }

    printf("Enter the elements in the second matrix :\n");
    for (i = 0; i < n; i++) {
        for (j = 0; j < n; j++) {
            printf("element - [%d],[%d] : ", i, j);
            scanf("%d", &brr1[i][j]);
        }
    }

    printf("\nThe First matrix is :\n");
    for (i = 0; i < n; i++) {
        printf("\n");
        for (j = 0; j < n; j++)
            printf("%d\t", arr1[i][j]);
    }

    printf("\nThe Second matrix is :\n");
    for (i = 0; i < n; i++) {
        printf("\n");
        for (j = 0; j < n; j++)
            printf("%d\t", brr1[i][j]);
    }


    for (i = 0; i < n; i++)     // Calculate the sum of the matrices
        for (j = 0; j < n; j++)
            crr1[i][j] = arr1[i][j] + brr1[i][j];

    printf("\nThe Addition of two matrix is : \n");
    for (i = 0; i < n; i++) {
        printf("\n");
        for (j = 0; j < n; j++)
            printf("%d\t", crr1[i][j]);
    }
}