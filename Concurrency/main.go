package main

import (
	"fmt"
	"runtime"
)

//var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARCH:\t", runtime.GOARCH)
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	//wg.Add(1)
	//go foo()
	foo()
	boo()

	//channelmain()
	//racemain()
	atomicmain()
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	//wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	//wg.Done()
}

func boo() {
	for i := 0; i < 10; i++ {
		fmt.Println("boo:", i)
	}
}
