package main

import (
	"fmt"
	"time"
)

const CorrectFormat = "2006/01/02"

func main() {
	var inputDate string
	fmt.Scan(&inputDate)

	if outputDate, err := time.Parse("2006.01.02", inputDate); err == nil {
		fmt.Println(outputDate.Format(CorrectFormat))
	} else if outputDate, err := time.Parse("2006-01-02", inputDate); err == nil {
		fmt.Println(outputDate.Format(CorrectFormat))
	} else if outputDate, err := time.Parse("2006/01/02", inputDate); err == nil {
		fmt.Println(outputDate.Format(CorrectFormat))
	} else {
		fmt.Println("Date format is invalid")
	}
}
