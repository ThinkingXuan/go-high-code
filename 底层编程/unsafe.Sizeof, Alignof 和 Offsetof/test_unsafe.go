package main

import (
	"fmt"
	"unsafe"
)

type struct1 struct {
	bool
	float64
	int16
} // 3 words 4words
type struct2 struct {
	float64
	int16
	bool
} // 2 words 3words
type struct3 struct {
	bool
	int16
	float64
} // 2 words 3words

func main() {
	fmt.Println(unsafe.Sizeof(struct1{}))
	fmt.Println(unsafe.Sizeof(struct2{}))
	fmt.Println(unsafe.Sizeof(struct3{}))
}
