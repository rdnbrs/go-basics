package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func ctxmain() {
	/*ctx := context.Background()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("------------------------")

	ctx, _ = context.WithCancel(ctx)
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)*/

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check:", ctx.Err())
	fmt.Println("go routine:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check:", ctx.Err())
	fmt.Println("go routine:", runtime.NumGoroutine())
	fmt.Println("exit")
	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("error check:", ctx.Err())
	fmt.Println("go routine:", runtime.NumGoroutine())
}
