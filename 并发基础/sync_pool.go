package main

import "sync"

func main() {
	// 新建一个池子
	bufferPool := &sync.Pool{
		New: func() interface{} {
			return 100
		}}
	// 从池中获取一个对象
	buffer := bufferPool.Get().(int)

	// 使用完后，放入池中
	bufferPool.Put(buffer)
}
