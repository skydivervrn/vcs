package main

import (
	"fmt"
	"time"
)

var dataChan = make(chan int, 1) // nolint:gochecknoglobals

func main() {
	go check() // DO NOT delete this line! it is used to check your code

	var timer = time.NewTimer(1 * time.Second)
	defer timer.Stop()

	// Write the missing code below!
	for {
		select {
		case data := <-dataChan:
			fmt.Println(data)
			timer.Reset(1 * time.Second)
		case <-timer.C:
			fmt.Println("exit")
			return
		}
	}
}
