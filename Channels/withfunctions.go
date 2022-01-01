package main

import (
	"fmt"
)

func withfunctionmain() {

	c := make(chan int)

	//send
	go foo(c)

	//receive
	boo(c)

	fmt.Println("Exited")
}

//send
func foo(c chan<- int) {
	c <- 15
}

//receive
func boo(c <-chan int) {
	fmt.Println(<-c)
}
