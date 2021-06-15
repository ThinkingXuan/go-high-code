package main

/*
union B1 {
	int i;
	float f;
};
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var b1 C.union_B1

	fmt.Printf("%T\n", b1)
	fmt.Println("b1.i:", *(*C.int)(unsafe.Pointer(&b1)))
	fmt.Println("b1.f:", *(*C.float)(unsafe.Pointer(&b1)))
}
