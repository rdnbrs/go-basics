package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr.No"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"Dr. No", 56}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
