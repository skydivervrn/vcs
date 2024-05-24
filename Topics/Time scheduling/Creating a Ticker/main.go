package main

import (
	"fmt"
	"time"
)

// Do not delete this block!
var (
	_ = fmt.Println
	_ = time.Time{}
)

func main() {
	// Create a Ticker
	ticker := time.NewTicker(500 * time.Millisecond)

	Check(ticker)
}
