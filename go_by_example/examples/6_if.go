package examples

import "fmt"

func IfStatements() {
	if 7%2 == 1 {
		fmt.Println("7 is not divided by 2")
	} else {
		fmt.Println("7 is divided by 2")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divided by 2")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 7 or 8 is divided by 2")
	}

	if num := 8; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num has one digital")
	} else {
		fmt.Println("num has many digitals")
	}
}
