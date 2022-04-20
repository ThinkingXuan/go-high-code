package main

import (
	"fmt"
	"sync"
)

//sync.map 之后，对 map 的读写，不需要加锁。并且它通过空间换时间的方式，使用 read 和 dirty 两个 map 来进行读写分离，降低锁时间来提高效率。

func main() {
	var m sync.Map

	// 写入
	m.Store("youxuan", 18)
	m.Store("youkai", 19)

	// 读取
	age, _ := m.Load("youxuan")
	fmt.Println(age)

	// 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)

		fmt.Println(name, age)
		return true
	})

	// 删除
	m.Delete("youxuan")

	age, ok := m.Load("youxuan")
	fmt.Println(age, ok)

	// 读取并写入
	m.LoadOrStore("youkefeng", 22)
	age, _ = m.Load("youkefeng")
	fmt.Println(age)

}
