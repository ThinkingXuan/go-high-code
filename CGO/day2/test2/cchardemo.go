package main

//#include <stdio.h>
import "C"

type CChar C.char

func (c *CChar) GoString() string {
	return C.GoString((*C.Char)(c))
}

func PrintCString(cs *C.Char) {
	C.puts(cs)
}

func main() {
}