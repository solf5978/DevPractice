#include <stdio.h>
int foo (int n)
{
    int i, m = 1;
    for(i = 1; i <= n; i++)
        m *= i;
    return m;
}

int main()
{
    int m, n;
    scanf("%d%d", &m, &n);
    printf("%d\n", foo(n)/(f(m) * f(n - m)));
    return 0;
}