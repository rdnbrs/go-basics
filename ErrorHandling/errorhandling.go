package main

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
}
