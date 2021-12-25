package main

import "fmt"

func main() {

	//standart for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	//while
	i := 15
	for i < 25 {

		i++
		if i == 17 {
			continue
		} else if i == 22 {
			fmt.Println("Breaking")
			break
		}
		fmt.Printf("%d\n", i)
	}

	//for { } infinit loop

	if x := 15; x == 15 {
		fmt.Printf("hahahahahahah %d\n", x)
	}

	y := 23
	switch y {
	case 21:
		fmt.Println("here 1")
	case 15:
		fmt.Println("here 2")
	default:
		fmt.Println("here default")
	}

}
