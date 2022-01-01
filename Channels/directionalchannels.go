package main

import (
	"fmt"
)

func directionalmain() {
	c := make(chan int, 2)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	c <- 44
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("---------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	fmt.Println("---------")
	cr = c
	cs = c

	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
