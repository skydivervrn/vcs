package main

import (
	"fmt"
)

func main() {
	var numberRange, targetCount, oddCount int
	fmt.Scan(&numberRange, &targetCount)

LoopBegin:
	for {
		number := numberRange

		for number > 0 {
			digit := number % 10
			number /= 10

			oddCount += digit % 2

			if oddCount >= targetCount {
				fmt.Println(numberRange)
				break LoopBegin
			}
		}

		numberRange++
	}
}
