package main

//void PrintString(_GoString_ s);
import "C"

import (
	"fmt"
)

func main() {
	C.PrintString("Hello, CGO\n")
}

//export PrintString
func PrintString(s string) {
	fmt.Print(s)
}