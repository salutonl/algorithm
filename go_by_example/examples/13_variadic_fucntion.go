package examples

import "fmt"

func Add(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Printf("sum: %v\n", sum)
	return sum
}
