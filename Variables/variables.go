package main

import "fmt"

type mvar int

const (
	a int    = 51
	c string = "baris"
)

func main() {
	x := 15
	var y string = "String val"
	var z = 15.5
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n%T\n%T\n", x, y, z)

	var b mvar = 42
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	x = int(b)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var t bool = true
	fmt.Println(t)
	fmt.Printf("%T\n", t)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%b\n", a)

	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
