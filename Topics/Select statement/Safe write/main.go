package main

import "fmt"

// Do not modify the `safeWrite` function name. It is used for testing
func safeWrite(inputData int) bool {
	// The channel is declared in the hidden `dataChan` variable
	// Write your code below:
	select {
	case dataChan <- inputData:
		return true
	default:
		return false
	}
}

// DO NOT modify the contents of the main function!
func main() {
	_ = fmt.Print
	checkSolution()
}
