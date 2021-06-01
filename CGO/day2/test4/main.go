package main

//char* cs = "hello";
import "C"
import "go_high_code/CGO/day2/test3"

func main() {
	test3.PrintCString(C.cs)
	print(C.GoString(C.cs))
}