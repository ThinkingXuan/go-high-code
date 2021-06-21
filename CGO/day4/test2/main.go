package main

import "C"

func main() {

}

//export number_add
func number_add(a, b C.int) C.int {
	return a+b
}
