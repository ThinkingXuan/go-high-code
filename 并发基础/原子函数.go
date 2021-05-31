package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var value uint64

func work1(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i < 100000; i++ {
		atomic.AddUint64(&value, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go work1(&wg)
	go work1(&wg)
	wg.Wait()

	fmt.Println(value)
}
