package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	sumAB := plus(1, 2)
	sumABC := plusPlus(3, 4, 5)
	fmt.Printf("sumAB: %v\n", sumAB)
	fmt.Printf("sumABC: %v\n", sumABC)

}