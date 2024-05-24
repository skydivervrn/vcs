package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type bankAccount struct {
	// add the struct tags below!
	AccountID   int     `json:"accountID"`
	AccountType string  `json:"accountType"`
	Balance     float64 `json:"balance"`
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var acc bankAccount
	fmt.Scanln(&acc.AccountID, &acc.AccountType, &acc.Balance)

	jsonResponse, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonResponse))
}
