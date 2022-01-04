#include <stdio.h>
#include <math.h>

int main()
{
    printf("%d\n", 11111 * 11111);
    printf("%d\n", 111111 * 111111);
    printf("%d\n", 111111111 * 111111111);
    printf("%lf\n", 1.1111 * 1.1111);
    printf("%d\n", sqrt(-10));
    printf("%d %d\n", 1.0/0.0, 0.0/0.0);
    printf("%d\n", 1/0);

    return 0;
}