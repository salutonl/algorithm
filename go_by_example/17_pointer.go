package main

import "fmt"

func zeroVal(n int) {
	n = 0
}

func zeroPrt(n *int) {
	*n = 0
}

func main() {
	a := 1

	zeroVal(a)
	fmt.Printf("a: %v\n", a)
	zeroPrt(&a)
	fmt.Printf("a: %v\n", a)
	fmt.Println(&a)

}