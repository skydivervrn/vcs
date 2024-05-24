package main

import "fmt"

func main() {
	numbers := [5]interface{}{65, 66, 67, 68, 69}
	for _, number := range numbers {
		letter := number.(int)
		fmt.Println(string(letter))
	}
}
