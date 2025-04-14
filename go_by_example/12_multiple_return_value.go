package main

import "fmt"

func val() (int, int) {
	return 3, 7
}

func main() {
	a, b := val()

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	_, c := val()
	fmt.Printf("c: %v\n", c)

}