package main

import "fmt"

func main() {
	var numberRange, end, sum3, sum5, other int
	fmt.Scan(&numberRange, &end)

	// put your code here
	for numberRange <= end {
		switch {
		case numberRange%3 == 0:
			sum3 += numberRange
		case numberRange%5 == 0:
			sum5 += numberRange
		default:
			other += numberRange
		}
		numberRange++
	}

	fmt.Println(sum3)
	fmt.Println(sum5)
	fmt.Println(other)
}
