package examples

import (
	"fmt"
	"time"
)

func SwitchStatements() {
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
}
