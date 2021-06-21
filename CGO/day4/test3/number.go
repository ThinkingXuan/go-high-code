package number

import "C"

//export number_add
func number_add(a, b C.int) C.int {
	return a+b
}
