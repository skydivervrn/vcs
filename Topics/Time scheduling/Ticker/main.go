package main

import (
	"fmt"
	"time"
)

// Do not delete this block!
var _ = fmt.Println

func MyTicker() *time.Ticker {
	return time.NewTicker(1 * time.Second)
}
