#include <stdio.h>

/* print Fahrenheit-Celsius table using For Loop */

int main()
{
    int fahr;
    //  Print in reversed order
    for (fahr = 300; fahr >= -18; fahr = fahr - 20)
        printf("%3d %6.1f\n", fahr, (5.0/9.0)*(fahr-32));
    return 0;
}