package main

import "C"

func main() {

}

//export get_string
func get_string() *C.char {
	return C.CString("HelloWorld")
}
