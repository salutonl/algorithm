package examples

import "fmt"

type IntError struct {
	arg int
	msg string
}

func (e *IntError) Error() string {
	return fmt.Sprintf("%d -- %s", e.arg, e.msg)
}

