package main

import "C"
import (
	"fmt"
	_ "go_high_code/CGO/day4/test4/number"
)

func main() {
	fmt.Println("Done")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln:", C.GoString(s))
}