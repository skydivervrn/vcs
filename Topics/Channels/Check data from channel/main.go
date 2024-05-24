package main

import "fmt"

const Attempts = 7

// set the condition in this function to stop the loop in time
func receiver(ch chan int) {
	for range [Attempts]struct{}{} {
		if value, ok := <-ch; ok {
			fmt.Println(value)
		}
	}
}

// do not change the code bellow
func sender(send int) chan int {
	ch := make(chan int, send)

	go func() {
		defer close(ch)
		for i := 0; i < send; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	var send int
	fmt.Scan(&send)

	ch := sender(send)
	receiver(ch)
}
