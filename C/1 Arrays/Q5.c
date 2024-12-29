// In this program, we will enter two sorted arrays. While merging them, we will compare the elements of both the arrays and merge them in a sorted manner.


#include <stdio.h>

int main()
{
    int n1,n2,n3;
    int a[400], b[400], c[400];
    printf("Enter the size of first array: ");
    scanf("%d",&n1);
    printf("Enter the array elements: ");
    for(int i = 0; i < n1; i++)      
       scanf("%d", &a[i]);
    printf("Enter the size of second array: ");
    scanf("%d",&n2);
    printf("Enter the array elements: ");
    for(int i = 0; i < n2; i++)      
       scanf("%d", &b[i]);
    n3 = n1 + n2;
    int i = 0, j = 0, k = 0;

    while (i < n1 && j < n2) {//this loop will run till a or b is completely traversed
        if (a[i] < b[j])
            c[k++] = a[i++];    //As soon as we copy an element in c, we increment the iterator so that the next element is copied at next index.
//When we copy an element from a to c, we increment i also because now we will compare with the next element of a.
        else
            c[k++] = b[j++];
    }
  
    while (i < n1)    //copying the leftover elements of a, if any
        c[k++] = a[i++];
  
    while (j < n2)    //copying the leftover elements of b, if any
        c[k++] = b[j++];
    
    printf("Final array after merging: ");
    for(int i = 0; i < n3 ; i++)
        printf(" %d ",c[i]);
    return 0;
}