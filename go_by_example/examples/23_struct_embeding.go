package examples

import "fmt"

type Base struct {
	Num int
}

func (b Base) Describe() string {
	return fmt.Sprintf("this is base %v", b.Num)
}

type Container struct {
	Base
	S string
}

type Describe interface {
	Describe() string
}