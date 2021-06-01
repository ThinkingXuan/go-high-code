package main

//#include "print_string.h"
import "C"

func main() {
	C.PrintString(C.CString("Hello CGO"))
}