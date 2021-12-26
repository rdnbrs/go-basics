package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type subobj struct {
	person
	ltk bool
}

func main() {
	fmt.Println("!!!Struct!!!")
	p1 := person{
		first: "Baris",
		last:  "Erden",
	}
	p2 := person{
		first: "ABC",
		last:  "DEF",
	}

	fmt.Println(p1.first)
	fmt.Println(p2)

	fmt.Println("!!!Struct with sub object!!!")
	s1 := subobj{
		ltk: true,
		person: person{
			first: "ABC",
			last:  "DEF",
		},
	}

	fmt.Println(s1)
	fmt.Println(s1.first, s1.last, s1.ltk)

	fmt.Println("!!!Anonymous Struct!!!")

	p3 := struct {
		first string
		last  string
	}{
		first: "bbbb",
		last:  "cccc",
	}

	fmt.Println(p3)

}
