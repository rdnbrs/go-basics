package main

import (
	"fmt"
	"sync"
)

func fanmain() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	//send
	go sendf(eve, odd)

	//receive
	go receivef(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func sendf(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receivef(e, o <-chan int, q chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			q <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			q <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(q)
}
