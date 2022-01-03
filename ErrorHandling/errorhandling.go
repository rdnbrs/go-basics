package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrGlobal = errors.New("Global Error")

func main() {
	/*
		//sample 1
		n, err := fmt.Println("Hello")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(n)
	*/

	/*
		//sample2
		var answer1, answer2, answer3 string
		fmt.Println("Name:")
		_, err := fmt.Scan(&answer1)
		if err != nil {
			panic(err)
		}

		fmt.Println("Fav:")
		_, err = fmt.Scan(&answer2)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sport:")
		_, err = fmt.Scan(&answer3)
		if err != nil {
			panic(err)
		}

		fmt.Println(answer1, answer2, answer3)
	*/
	/*
		//sample3
		f, err := os.Create("names.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		r := strings.NewReader("asjdfhaklsjdfhaskjdhf")
		io.Copy(f, r)
	*/

	/*
		//sample4
		f, err := os.Open("names.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bs))
	*/
	/*
		//sample5
		_, err := os.Open("nofile.txt")
		if err != nil {
			log.Println("err happened", err)
		}
	*/
	/*
		//sample 6
		f, err := os.Create("names.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		log.SetOutput(f)

		f2, err := os.Open("asdasd.txt")
		if err != nil {
			log.Println("err happened", err)
		}
		defer f2.Close()
	*/

	//Deffer
	var x int
	x++
	fmt.Println(x)
	i := c()
	fmt.Println(i)

	//recover
	f()
	fmt.Println("Returned normally from f")

	fmt.Println("!!!ERRORS WITH INFO!!!")

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 5
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("Calling g")
	g(0)
	fmt.Println("returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func sqrt(i float64) (float64, error) {
	if i < 0 {
		//return 0, errors.New("square root of negative number")
		//return 0, fmt.Errorf("%v", i)
		return 0, ErrGlobal
	}

	return i * i, nil
}
