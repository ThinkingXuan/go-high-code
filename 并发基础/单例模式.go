package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

// 基于 sync.Once实现
var (
	once sync.Once
)

func Instance2() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
