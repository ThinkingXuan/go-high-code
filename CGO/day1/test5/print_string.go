package main

import "C"
import "fmt"

// 将Go程序代码导出为C语言程序，使用export关键字，//export之间不能有空格

//export PrintString
func PrintString(s *C.char) {
	fmt.Print(C.GoString(s))
}
