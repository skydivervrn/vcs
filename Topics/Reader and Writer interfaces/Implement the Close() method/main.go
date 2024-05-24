package main

import (
	"fmt"
	"io"
)

type FunnyBox struct{}

// Create the `Close()` method for the `FunnyBox` struct below:
func (b FunnyBox) Close() error {
	fmt.Println("Bey!")
	return nil
}

// DO NOT delete or modify the code within the main() function!
// nolint: gosimple // DO NOT delete this comment!
func main() {
	var c io.Closer
	c = &FunnyBox{}
	c.Close()
}
