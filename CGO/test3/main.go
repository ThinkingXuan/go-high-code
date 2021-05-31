package main

/*
#include <stdio.h>
static void PrintString(char* str) {
	puts(str);
}
 */
import "C"

func main() {
	C.PrintString(C.CString("Hello CGO"))
}
