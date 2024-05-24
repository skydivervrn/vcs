package main

import (
	"fmt"
	"math/rand"
	"time"
)

// complete these functions
func mysteryOne(data []int, ch <-chan int) {
	for val := range ch {
		fmt.Print(val)
	}
}

func mysteryTwo(data []int, ch chan<- int) {
	for _, item := range data {
		ch <- item
	}
}

// DO NOT modify the code below!
func main() {
	_ = rand.Int() // DO NOT delete this `rand.Int()` line!

	var length int
	fmt.Scanln(&length)
	numbers := randomizer(length)
	ch := make(chan int)

	go mysteryOne(numbers, ch)
	go mysteryTwo(numbers, ch)

	time.Sleep(time.Second)
}
