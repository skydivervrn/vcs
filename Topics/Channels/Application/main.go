package main

import "fmt"

func sender(ch chan int) {
	ch <- 1
}

func receiver(ch chan int) {
	number := <-ch
	fmt.Print(number)
}

// do not change code below
func main() {
	var amount int
	fmt.Scan(&amount)

	ch := make(chan int)
	for i := 0; i < amount; i++ {
		go sender(ch)
	}

	for i := 0; i < amount; i++ {
		receiver(ch)
	}
}
