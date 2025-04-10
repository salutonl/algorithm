package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["first"] = 1
	m["second"] = 2
	fmt.Println(m)

	m1 := m["first"]
	m2 := m["third"]

	fmt.Println(m1, m2)

	if _, prs := m["third"]; prs == true {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}

	fmt.Println(len(m))

	delete(m, "second")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	m3 := map[string]int{"test": 1, "test2": 2}
	fmt.Println(m3)
	m4 := map[string]int{
		"test":  1,
		"test2": 2,
	}
	fmt.Println(m4)
	fmt.Println(maps.Equal(m3, m4))
}
