package test3

//#include <stdio.h>
import "C"

func PrintCString(cs *C.char) {
	C.puts(cs)
}