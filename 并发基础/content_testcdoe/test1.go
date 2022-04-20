package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go read(ctx, "1")
	go read(ctx, "2")
	go read(ctx, "3")

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

}

func read(ctx context.Context, name string) {

	go read2(ctx, "4")
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " 已终止")
			return
		default:
			fmt.Println(name, "监控中")
			time.Sleep(2 * time.Second)
		}
	}
}

func read2(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " 已终止")
			return
		default:
			fmt.Println(name, "监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
