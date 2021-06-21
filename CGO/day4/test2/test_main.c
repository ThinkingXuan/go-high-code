#include "number.h"
#include <stdio.h>

int main() {
    int a = 10;
    int b = 5;

    int c = number_add(a,b);

    printf("%d+%d = %d\n", a, b ,c);
    return 0;
}
