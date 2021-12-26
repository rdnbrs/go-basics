package main

import (
	"fmt"
)

func passbyvalue(b int) {
	b = 5
}

func passbyreference(b *int) {
	*b = 8
}

func main() {
	a := 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//address
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)

	passbyvalue(a)
	fmt.Println(a)
	passbyreference(&a)
	fmt.Println(a)
}
