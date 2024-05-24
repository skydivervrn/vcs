package main

import "fmt"

func main() {
	var i, num, sum int
	fmt.Scan(&num)

SomeLabel:
	if i%2 == 0 {
		sum += i
	}
	i++
	if i <= num {
		goto SomeLabel
	}

	fmt.Println(sum)
}
