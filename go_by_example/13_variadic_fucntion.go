package main

import (
	"fmt"
)

func add(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Printf("sum: %v\n", sum)
	return sum
}

func main() {
	add(1,2,3,4,5)
	add(6,7,8,9)
	slice := []int{10,9,8}
	fmt.Printf("%v\n", add(slice...))
}