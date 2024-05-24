package main

func main() {
	// you need to change the code only in this place
	ch := make(chan struct{}, 5)

	// do not change the code below
	for range [5]struct{}{} {
		ch <- struct{}{}
	}

	secret(ch)
}
