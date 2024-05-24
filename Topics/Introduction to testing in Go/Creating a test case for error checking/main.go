package main

import (
	"errors"
	"fmt"
	"testing"
)

var ErrDivByZero = errors.New("division by zero")
var Div func(a, b int) (int, error) // nolint: gochecknoglobals // DO NOT delete this comment!

func TestDivByZero(t *testing.T) {
	// replace all "_"
	_, err := Div(10, 0)
	if err != nil && !errors.Is(err, ErrDivByZero) {
		t.Errorf("got != expected: %v", err)
	}
}

// DO NOT MODIFY
func main() {
	var t testing.T

	Div = func(a, b int) (int, error) {
		if b == 0 {
			return 0, ErrDivByZero
		}
		fmt.Println("b != 0: in this task you only need to check the case when b == 0")
		return 0, nil
	}

	TestDivByZero(&t)

	fmt.Println("OK")
}
