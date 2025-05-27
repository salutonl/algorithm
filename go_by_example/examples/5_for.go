package examples

import "fmt"

func ForLoop() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 6 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for k := range 6 {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
