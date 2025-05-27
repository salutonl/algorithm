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


}
