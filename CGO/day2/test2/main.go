package main

//char* cs = "hello";
import "C"

func main() {
	PrintCString(C.cs)
	print(C.GoString(C.cs))
}