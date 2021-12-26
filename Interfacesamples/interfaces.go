package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

type secret struct {
	person
	ltk bool
}

func bar(h human) {
	h.speak()
}

func (s secret) speak() {
	fmt.Println(s)
}

func (s person) speak() {
	fmt.Println(s)
}

func main() {
	s1 := secret{
		person: person{
			"baris",
			"erden",
		},
		ltk: false,
	}

	bar(s1)
	bar(s1.person)
}
