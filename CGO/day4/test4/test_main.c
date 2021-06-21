#include <stdio.h>

void goPrintln(char*);
int number_add(int a, int b);

int main() {
    int a = 10;
    int b = 5;

    int c = number_add(a, b);
    printf("(%d+%d) = %d\n", a, b, c);

    goPrintln("done");
    return 0;
}