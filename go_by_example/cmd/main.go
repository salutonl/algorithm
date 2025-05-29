package main

import (
	"algorithm/go_by_example/examples"
	"fmt"
)

func main() {
	fmt.Println("Running examples:")

	// Call example functions
	examples.HelloWorld()
	examples.Arrays()
	examples.Maps()
	_, a := examples.Vals()
	fmt.Println(a)
	examples.Range()

	p := examples.NewPerson("Har")

	fmt.Printf("%T\n", p)
	fmt.Println(examples.GetPersonAge(p))
	dog := struct{
		name string
		age int
	}{name: "dou", age: 14}

	fmt.Println(dog)

	r := examples.Rect{10, 5}

	fmt.Println(r.Area())
	fmt.Println(r.Perim())

	rp := &examples.Rect{Height:2, Width:3}
	fmt.Println(rp.Perim())
	fmt.Println(rp.Area())

	fang := &examples.Fang{2.1, 3.2}
	yuan := &examples.Yuan{1.2}

	examples.Measure(fang)
	examples.Measure(yuan)

}
