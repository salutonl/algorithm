package examples

import (
	"fmt"
	"math"
)

type Geometry interface {
	Xy() float64
	XPlusY() float64
}

type Yuan struct {
	Jing float64
}

type Fang struct {
	Chang, Kuan float64
}

func (r *Yuan) Xy() float64 {
	return math.Pi * r.Jing * r.Jing
}

func (r *Yuan) XPlusY() float64 {
	return 2 * math.Pi * r.Jing
}

func (r *Fang) Xy() float64 {
	return r.Chang * r.Kuan
}

func (r *Fang) XPlusY() float64 {
	return 2 * (r.Chang+r.Kuan)
}

func Measure(g Geometry) {
	if _, ok := g.(*Yuan); ok {
		fmt.Println("This is a Yuan")
	} else if _, ok := g.(*Fang); ok {
		fmt.Println("This is a Fang")
	}
	xy := g.Xy()
	fmt.Println(xy)

	xPlusY := g.XPlusY()
	fmt.Println(xPlusY)
}
