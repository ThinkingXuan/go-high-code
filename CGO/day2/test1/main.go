package main

/*
#include<stdio.h>
void printInt(int v) {
	printf("printInt,: %d\n",v);
}
 */
import "C"

func main()  {
	v := 100
	C.printInt(C.int(v))
}
