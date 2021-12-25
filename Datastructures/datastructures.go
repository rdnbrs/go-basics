package main

import "fmt"

func main() {

	//go documentation prefer use slices rather than arrays
	//array
	fmt.Println("!!!Arrays!!!")
	var x [5]int
	fmt.Println(x)
	x[3] = 13
	fmt.Println(x)
	fmt.Println(len(x))

	//slice
	//t := type{values}
	fmt.Println("!!!Slices!!!")
	t := []int{4, 5, 7, 8, 11}
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Printf("%T\n", t)

	fmt.Println("!!!Slice range!!!")
	for i, v := range t {
		fmt.Println(i, v)
	}

	fmt.Println("!!!Slice a slice!!!")
	fmt.Println(t[0:1])
	fmt.Println(t[2:])
	fmt.Println(t[:3])

	fmt.Println("!!!Slice append!!!")
	t = append(t, 44, 22)
	fmt.Println(t)
	fmt.Println(len(t))
	y := []int{1, 2, 3}
	t = append(t, y...)
	fmt.Println(t)
	fmt.Println(len(t))

	fmt.Println("!!!Slice delete item!!!")
	//remove t[2] and t[3]
	t = append(t[:2], t[4:]...)
	fmt.Println(t)
	fmt.Println(len(t))

	fmt.Println("!!!Slice make!!!")
	d := make([]int, 10, 12)
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	d[0] = 1
	d[9] = 999
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	fmt.Println("!!!Multi-dimensional slice!!!")
	md := []string{"a", "b", "c"}
	fmt.Println(md)
	md2 := []string{"d", "e", "f"}
	fmt.Println(md2)
	xp := [][]string{md, md2}
	fmt.Println(xp)

	fmt.Println("!!!Maps!!!")
	m := map[string]int{
		"key1": 15,
		"key2": 25,
	}
	fmt.Println(m)
	fmt.Println(m["key1"])
	if v, ok := m["keys3"]; ok {
		fmt.Println(v)
	}

	fmt.Println("!!!Maps add item!!!")
	m["key3"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("!!!Maps delete key!!!")
	delete(m, "key2")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
