package main

// 声明C语言的函数，相当于 .h 文件

//void PrintString(char* s);
import "C"

import "fmt"


func main() {
	C.PrintString(C.CString("Hello, CGO\n"))
}

//export PrintString
func PrintString(s *C.char) {
	fmt.Print(C.GoString(s))
}
