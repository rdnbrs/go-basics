package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	bs2 := []byte(string(bs))

	var unpeople []person

	err2 := json.Unmarshal(bs2, &unpeople)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(unpeople)

}
