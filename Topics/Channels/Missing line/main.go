package main

import (
	"fmt"
	"time"
)

// fix this function by adding the missing line
func sender(ch chan int, amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(200 * time.Millisecond)
		ch <- 1
	}
	close(ch)
}

// do not change the code bellow
func main() {
	var amount int
	fmt.Scan(&amount)

	ch := make(chan int)
	go sender(ch, amount)

	for range ch {
		fmt.Println("done")
	}
}
