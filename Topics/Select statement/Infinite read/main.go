package main

import (
	"fmt"
)

var chan1 = make(chan int, 1) // nolint:gochecknoglobals
var chan2 = make(chan int, 1) // nolint:gochecknoglobals
var chan3 = make(chan int, 1) // nolint:gochecknoglobals

func main() {
	_ = fmt.Print
	check() // DO NOT delete this line! it is used to check your code
	var sum int

	// Write the missing code below!
	for {
		select {
		case num := <-chan1:
			sum += num
		case num := <-chan2:
			sum += num
		case num := <-chan3:
			sum += num
		default:
			fmt.Println(sum)
			return
		}
	}
}
