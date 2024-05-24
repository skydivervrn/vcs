package main

import (
	"fmt"
	"time"
)

// complete this function
func sender(data []int) <-chan int {
	ch := make(chan int)
	ch2 := (<-chan int)(ch)
	go func() {
		defer func() {
			close(ch)
		}()
		for _, item := range data {
			ch <- item
		}
	}()

	return ch2
}

// DO NOT change the code bellow
func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Print(val)
	}
}

func main() {
	var length int
	fmt.Scanln(&length)
	data := randomizer(length)

	ch := sender(data)

	if !isReceiver(ch) {
		fmt.Println("channel should be read-only")
		return
	}
	go receiver(ch)

	time.Sleep(time.Second)
}
