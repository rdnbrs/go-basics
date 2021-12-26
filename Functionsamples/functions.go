package main

import (
	"fmt"
)

//everything in go is pass by value
func main() {
	fmt.Println("!!!Basics!!!")
	foo()
	var x int = foo2()
	fmt.Println(x)
	bar("asd")
	a, b := foo3()
	fmt.Println(a, b)
	foo4(1, 2, 3, 4, 5, 6)

	xi := []int{1, 2, 3, 4, 5, 6}
	foo4(xi...)

	fmt.Println("!!!Defer!!!")
	foo()
	bar("asd")
	defer foo()
	bar("asd")
	bar("asd")
}

//basic
func foo() {
	fmt.Println("foo void function")
}

//one return
func foo2() int {
	fmt.Println("foo2 return int function")
	return 5
}

//with argument
func bar(s string) {
	fmt.Println("Hello", s)
}

//two return
func foo3() (int, string) {
	return 5, "asd"
}

//unlimited argument (slice)
func foo4(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var sum int = 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}
