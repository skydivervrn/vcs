package main

import "fmt"

var inputChan = make(chan int, 1)  // nolint:gochecknoglobals
var outputChan = make(chan int, 1) // nolint:gochecknoglobals

func main() {
	check() // DO NOT delete this line! it is used to check your code

	// Write the missing code below!
	for {
		select {
		case input := <-inputChan:
			fmt.Println(input)
		case outputChan <- 5:
		default:
			fmt.Println("exit")
			return
		}
	}
}
