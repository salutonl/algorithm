package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 400
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println(len(a))


	b := [6]int{1,2,2,3,4}
	fmt.Println(b)

	c := [...]string{"a", "bb", "cc"}
	fmt.Println(c)

	d := [...]int{100, 200, 4:300, 400}
	fmt.Println(d)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	threeD := [2][3]int{
		{1,2,3},
		{2,3,4},
	}
	fmt.Println(threeD)
}