package examples

import (
	"fmt"
)

const a = "apple"

func Constants() {
	fmt.Println(a)

	const b = 5000000
	fmt.Println(b)

	const c = 6e20
	fmt.Println(c)

	const d = 6e20 / b
	fmt.Println(d)

	fmt.Println(int32(b))
	fmt.Println(int64(d))
}
