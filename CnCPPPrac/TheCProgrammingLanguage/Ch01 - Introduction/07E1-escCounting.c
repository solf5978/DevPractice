#include <stdio.h>

/* Count Lines in input */

int main()
{
    int c, ns = 0, nt = 0, nl = 0;

    while((c = getchar()) != EOF)
        if (c == ' ')
            ++ns;
        if (c == '\t')
            ++nt;
        if (c == '\n')
            ++nl;
        printf("Space: %d, Tab: %d, NewLine: %d\n", ns, nt, nl);
}