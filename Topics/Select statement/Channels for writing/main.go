package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan int, 1) // nolint:gochecknoglobals
var chan2 = make(chan int, 1) // nolint:gochecknoglobals

func main() {
	var inputData int

	go func() { // DO NOT delete this line!
		// Write the missing code below!
		for {
			fmt.Scan(&inputData)

			select {
			case chan1 <- inputData:
			case chan2 <- inputData:
			}
		}
	}()

	_ = time.Now
	check() // DO NOT delete this line! it is used to check your code
}
