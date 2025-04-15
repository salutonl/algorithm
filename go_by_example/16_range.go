package main

import "fmt"


func main() {

	for i := range 5 {
		fmt.Printf("i: %v\n", i)
	}

	s := []int{1,2,3,4,5,6}

	for i, v := range s {
		fmt.Println(i, v)
	}

	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("sum: %v\n", sum)

	m := map[string]string{"first": "banana", "second": "apple"}

	for k, v := range m {
		println(k, v)
	}

	for k := range m {
		println(k)
	}

	for i, c := range "abcdef" {
		println(i, c)
	}
}