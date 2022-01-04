#include <stdio.h>

/* Print Fahrenheit-Celsius table
 * for fahr = 0, 20, ..., 300
*/

int main()
{
    float fahr, celsius;
    int lower, upper, step;

    lower = -17;      // Lower limit of temperature table
    upper = 148;    // upper limit
    step = 12;      // step size

    celsius = lower;
    while (celsius <= upper) {
        fahr = celsius * (9.0 / 5.0) + 32;
        //celsius = 5 * (fahr - 32) / 9;
        //Or simply change fahr and celsius
        // printf("%d\t %d\n", fahr, celsius);
        printf("%3.1f %6.0f\n", celsius, fahr);
        celsius += step;
    }
    return 0;
}