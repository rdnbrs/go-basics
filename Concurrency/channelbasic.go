package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func channelmain() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()

	fmt.Println(<-ch)
}
