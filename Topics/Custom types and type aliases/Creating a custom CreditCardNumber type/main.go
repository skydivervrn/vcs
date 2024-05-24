package main

import "fmt"

// Declare a custom type `CreditCardNumber` derived from `string`:
type CreditCardNumber string

func main() {
	// Declare a variable `ccn` of the custom `CreditCardNumber` type:
	var ccn CreditCardNumber
	fmt.Scanln(&ccn)

	// Check if the credit card number `ccn` is valid using the `IsValid()` method:
	if ccn.IsValid() {
		fmt.Println("This is a valid Credit Card Number!")
	} else {
		fmt.Println("This is not a valid Credit Card Number!")
	}
}

// DO NOT delete or modify the `IsValid()` method below:
// nolint: gomnd // DO NOT remove this comment!
func (ccn CreditCardNumber) IsValid() bool {
	var sum int
	reverseOrder := false
	for i := len(ccn) - 1; i >= 0; i-- {
		n := int(ccn[i] - '0')
		if reverseOrder {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		reverseOrder = !reverseOrder
	}
	return sum%10 == 0
}
