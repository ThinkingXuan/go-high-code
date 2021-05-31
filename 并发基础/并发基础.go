package main

import (
	"fmt"
	"sync"
)

// 使用Sync.Mutex互斥锁来实现并发 相加

var  total struct {
	sync.Mutex
	value int
}


func work(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		total.Lock()
		total.value += i
		total.Unlock()

		//atomic.AddInt64(&total.value, i)
	}
}

func main() {
	var wg sync.WaitGroup
	
	wg.Add(2)
	go work(&wg)
	go work(&wg)

	wg.Wait()

	fmt.Println(total.value)
}
