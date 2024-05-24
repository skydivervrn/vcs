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

	result := [5]int{}
	for i := 0; i < 5; i++ {
		// run the code in the loop concurrently
		go func(i int) {
			value := cube(i)
			// add the value to the result array
			result[i] = value
		}(i)
	}
	time.Sleep(time.Second)

	printResult(result)
}
