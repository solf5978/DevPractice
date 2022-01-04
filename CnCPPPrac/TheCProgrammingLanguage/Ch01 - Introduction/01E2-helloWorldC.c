#include <stdio.h>

/* 
 * Experiment to find out what happens when printf's argument string contains \c, 
 * where c is some character not listed above.
 * Ran But Warned: Unknown escape sequence: '\c' 
*/

int main()
{
    printf("Hello, World\c");
}