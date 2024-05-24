package main

import (
	"fmt"
)

// make it great famous brothers
func main() {
	ch := make(chan string, 2)

	luigi(ch)
	mario(ch)
	//runtime.Gosched()
}

// DO NOT change brother's functions
func luigi(ch chan<- string) {
	ch <- "Luigi"
	ch <- "Mario"
}

func mario(ch <-chan string) {
	fmt.Printf("Hello, %s Mario!\n", <-ch)
	fmt.Printf("Hello, %s Mario!\n", <-ch)
}
