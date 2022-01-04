#include <stdio.h>

/* copy input to output */

int main()
{
    int c;
    while (1){
        if (c = getchar() != EOF) {
            putchar(c);
            printf("%d\n", getchar() != EOF);
        }
        printf("%d\n", EOF);
    }
    
    return 0;
}