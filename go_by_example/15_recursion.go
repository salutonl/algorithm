package main

import "fmt"

func fact(n int) int {
	if n == 0{
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	fmt.Printf("fact(5): %v\n", fact(5))


	var fib func(n int) int 
	
	fib = func(n int) int {
		if n < 2 {
			return 1
		}
		return fib(n - 1) + fib(n - 2)
	}

	fmt.Printf("fib(4): %v\n", fib(4))

}