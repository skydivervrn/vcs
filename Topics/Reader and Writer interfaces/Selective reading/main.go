package main

import (
	"fmt"
	"io"
	"strings"
)

const textSource = "Supercalifragilisticexpialidocious"

// nolint: gomnd // DO NOT delete this comment!
func main() {
	fmt.Println("full text:", textSource)

	reader := strings.NewReader(textSource)

	p := make([]byte, 7)
	// Move the position of the `reader` to the last 7 bytes below:
	_, err := reader.Seek(int64(len(textSource)-7), io.SeekStart)
	if err != nil {
		return
	}

	if _, err := reader.Read(p); err != nil {
		fmt.Println(err)
	}
	fmt.Print("selected: ", string(p), " ")
}
