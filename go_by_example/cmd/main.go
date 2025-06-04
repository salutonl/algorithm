package main

import (
	"algorithm/go_by_example/examples"
	"errors"
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
	
	enums := examples.Trans(examples.ServiceConnected)
	fmt.Println(enums)

	enums2 := examples.Trans(examples.ServiceIdle)
	fmt.Println(enums2)

	embed := examples.Container{
		Base: examples.Base{Num:2},
		S: "Peter",
	}
	fmt.Printf("this is embed Num: %v, S:%v and base.num %v\n", embed.Num, embed.S, embed.Base.Num)
	fmt.Println(embed.Describe())

	var de examples.Describe = embed
	fmt.Println(de.Describe())

	index, value := examples.SliceIndex([]string{"apple", "banana", "orange"}, "banana")

	if index != -1 {
		fmt.Printf("index is %v, value is %v \n", index, value)
	} else {
		fmt.Printf("no match \n")
	}

	l := &examples.List[int]{}

	fmt.Println(l.AllElements())

	l.Push(3)
	l.Push(2)

	fmt.Println(l.AllElements())

	for i := range l.ShowAll() {
		fmt.Printf("list %v by iterators\n", i)
	}

	s := &examples.List[string]{}
	s.Push("hahhah")
	s.Push("ababab")
	s.Push("agohahah")

	for j := range s.ShowAll() {
		fmt.Printf("a string iterators %v \n", j)
	}

	for k := range examples.GetFib() {
		if k > 40 {
			break
		}
		fmt.Printf("A Fib iterator %v \n", k)
	}

	if _, e := examples.F(100); e != nil {
		fmt.Println("failed ", e)
	} else {
		fmt.Println("worked")
	}


	for i := range 5 {
		fmt.Printf("range for %v \n", i)
		e := examples.MakeTea(i)
		if errors.Is(e, examples.ErrNoPower) {
			fmt.Println("we should generate power")
		} else if errors.Is(e, examples.ErrNoTea){
			fmt.Println("go by some tea")
		} else if e != nil {
			fmt.Println("unknown err")
		} else {
			fmt.Println("tea is ok")
		}
	}
}
