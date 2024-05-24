package main

import (
	"fmt"
	"time"
)

// Do not delete this block!
var _ = fmt.Println

func Defuse(bomb *time.Timer) bool {
	// Adams %#@
	bomb.Stop()
	return true
}
