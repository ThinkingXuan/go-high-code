package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type struct1 struct {
	b bool
	f float64
	i int16
}
type struct2 struct {
	f float64
	i int16
	b bool
}

func main() {
	//在struct中，它的对齐值是它的成员中的最大对齐值。
	s1 := struct1{
		b: false,
		f: 9.9,
		i: 10,
	}
	s2 := struct2{
		f: 9.9,
		i: 10,
		b: false,
	}

	// 结构体1，
	fmt.Printf("%v => %v, %v, %v\n", unsafe.Alignof(s1), unsafe.Alignof(s1.b), unsafe.Alignof(s1.f), unsafe.Alignof(s1.i))
	fmt.Printf("%v => %v, %v, %v\n", reflect.TypeOf(s1).Align(), reflect.TypeOf(s1.b).FieldAlign(), reflect.TypeOf(s1.f).FieldAlign(), reflect.TypeOf(s1.i).FieldAlign())

	// 结构体2，
	fmt.Printf("%v => %v, %v, %v\n", unsafe.Alignof(s2), unsafe.Alignof(s2.f), unsafe.Alignof(s2.i), unsafe.Alignof(s2.b))
	fmt.Printf("%v => %v, %v, %v\n", reflect.TypeOf(s2).Align(), reflect.TypeOf(s2.f).FieldAlign(), reflect.TypeOf(s2.i).FieldAlign(), reflect.TypeOf(s2.b).FieldAlign())

	fmt.Println(unsafe.Sizeof(s1))
	fmt.Println(unsafe.Sizeof(s2))

}
