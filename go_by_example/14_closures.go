package main

import "fmt"


func counter() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}


func main() {
	c := counter()

	fmt.Printf("c(): %v\n", c())
	fmt.Printf("c(): %v\n", c())
	fmt.Printf("c(): %v\n", c())

	d := counter()

	fmt.Printf("d(): %v\n", d())
	fmt.Printf("d(): %v\n", d())
}