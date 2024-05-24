package main

import (
	"fmt"
)

// Create a type alias LightLevel based on int:
type LightLevel = int

// A function that answers the question "how dark it is" based on the LightLevel provided
// nolint: gomnd // DO NOT remove this comment!
func howDark(ll LightLevel) string {
	// Your code here
	if ll < 10 {
		return "It's very dark and scary!"
	}
	if ll < 25 {
		return "It's dark and not scary."
	}
	if ll < 50 {
		return "It's a bit dark in here, don't you think?"
	}
	if ll < 75 {
		return "There is plenty of light in here."
	}
	return "So bright! I better put on those sunglasses."
}

func main() {
	// DO NOT change the code below
	var lightLevel int
	fmt.Scan(&lightLevel)

	// use howDark function with provided light level to know how dark it is
	fmt.Println(howDark(lightLevel))
}
