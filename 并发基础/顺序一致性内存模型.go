package main

import "fmt"

var a string
var done bool

func setup() {
	a = "hello world"
	done = true
}

func main() {
	go setup()
	for !done {}
	//Go语言并不保证在main函数中观测到的对done的写入操作发生在对字符串a的写入的操作之后，因此程序很可能打印一个空字符串。
	//更糟糕的是，因为两个线程之间没有同步事件，setup线程对done的写入操作甚至无法被main线程看到，main函数有可能陷入死循环中。
	fmt.Println(a)
}
