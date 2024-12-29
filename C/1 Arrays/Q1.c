// This program reads a user-defined number of integer values into an array and then display these values in reverse order. 
// After storing the values, the program first prints them in the original order and then prints them in the reversed order.

#include <stdio.h>

int main()
{
   int i, n, a[400];

   printf("Input the number of elements to store in the array :");
   scanf("%d", &n);
   printf("Input %d elements in the array :\n", n);
   
   for (i = 0; i < n; i++) {
      printf("element - %d : ", i +1);
      scanf("%d", &a[i]);  
   }

   printf("\nThe values stored in the array are : \n");
   for (i = 0; i < n; i++) {
       printf("% 5d", a[i]);  
   }
   
   printf("\n\nThe values stored in the array in reverse are :\n");
   for (i = n - 1; i >= 0; i--) {
       printf("% 5d", a[i]);
   }

   return 0;
}