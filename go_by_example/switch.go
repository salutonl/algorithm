package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("i is one")
	case 2:
		fmt.Println("i is two")
	case 3:
		fmt.Println("i is three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	default:
		fmt.Println("afternoon")
	}

	whatAmI := func (i interface{})  {
		switch i.(type) {
		case int:
			fmt.Println("its number")
		case bool:
			fmt.Println("its bool")
		default:
			fmt.Println("dont know")
		}

	}
	whatAmI("apple")
	whatAmI(1)
	whatAmI(true)
}