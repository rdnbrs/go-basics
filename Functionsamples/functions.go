package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secret struct {
	person
	ltk bool
}

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

	/*defer runs end of the main
	fmt.Println("!!!Defer!!!")
	foo()
	bar("asd")
	defer foo()
	bar("asd")
	bar("asd")*/

	s1 := secret{
		person: person{
			first: "abc",
			last:  "sdsds",
		},
		ltk: true,
	}

	printStruct(s1)

	f1 := func() {
		fmt.Println("!!!Anonymous function!!!")
	}

	f1()

	func(x int) {
		fmt.Println(x)
	}(42)

	//function return from function
	f2 := ffunc()
	fmt.Println(f2())

	//callback
	fmt.Println("!!!Callback function!!!")
	fmt.Println(even(sum, xi...))

	fmt.Println("!!!Scope!!!")
	f3 := incrementor()
	f4 := incrementor()
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f4())

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

//struct as argument
func printStruct(s secret) {
	fmt.Println(s)
}

//return func
func ffunc() func() int {
	return func() int {
		return 1984
	}
}

//callback func
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	var response int = f(yi...)
	return response
}

//scope
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
