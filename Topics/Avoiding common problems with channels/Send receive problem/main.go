package main

import (
	"fmt"
	"time"
)

const Hello = "Hello, "

// fix the function argument
func receiver(ch <-chan string) {
	if isCommon(ch) {
		fmt.Println("cheating detected!")
		return
	}

	for val := range ch {
		fmt.Print(val)
	}
}

// fix the function argument
func sender(name string, ch chan<- string) {
	if isCommon(ch) {
		fmt.Println("cheating detected!")
		return
	}

	ch <- Hello
	ch <- name
	close(ch)
}

// DO NOT change yhe code bellow
func main() {
	var name string
	fmt.Scanln(&name)
	ch := make(chan string)

	go receiver(ch)
	go sender(name, ch)

	time.Sleep(time.Second)
}
