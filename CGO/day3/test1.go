package main

/*
struct A {
	int i;
	float f;
};
struct B {
	int type;        // type是go语言的关键字
};
struct C {
	int type;        // type是go语言的关键字
	int _type;        // type和_type同时存在时，type关键字被屏蔽
};
 */
import "C"
import "fmt"

func main()  {
	var a C.struct_A
	a.i = 10
	a.f = 9.9
	fmt.Println(a.i)
	fmt.Println(a.f)

	var b C.struct_B
	//b.type = 11         编译错误
	b._type = 11
	fmt.Println(b._type)
}
