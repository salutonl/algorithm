package examples

import (
	// "errors"
	"fmt"
)

func F(arg int) (int, error) {
	if arg > 42 {
		return arg, &IntError{100, "too old la"}
	}

	return arg, nil
}

var ErrNoTea = fmt.Errorf("we should buy some tea")
var ErrNoPower = fmt.Errorf("we should generate power")

func MakeTea(arg int) error {
	if arg == 2 {
		return ErrNoTea
	} else if arg == 4{
		return fmt.Errorf("embed %w", ErrNoPower)
	}
	return nil
}


