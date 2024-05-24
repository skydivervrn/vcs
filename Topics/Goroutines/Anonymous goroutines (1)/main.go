package main

import (
	"fmt"
	"time"
)

func main() {
	// do not modify this block of code
	start := time.Now()
	defer func() {
		if time.Since(start) > time.Second*2 {
			fmt.Println(10000)
		}
	}()

	for i := 0; i < 5; i++ {
		// run the code in the loop concurrently
		go func(i int) {
			// this simulates long calculations
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("%d ", i*i)
		}(i)
	}
	time.Sleep(time.Second)

}
