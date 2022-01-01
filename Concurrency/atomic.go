package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func atomicmain() {
	fmt.Println("!!!Race Condition!!!")

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println(counter)
}