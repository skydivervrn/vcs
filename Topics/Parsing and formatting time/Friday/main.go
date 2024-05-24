package main

import (
	"fmt"
	"log"
	"time"
)

const (
	HoursPerDay    = 24
	MinutesPerHour = 60
)

// DO NOT modify the getFridayDate function:
func getFridayDate(currentDate time.Time) time.Time {
	fridayDate := time.Date(
		currentDate.Year(),
		currentDate.Month(),
		currentDate.Day(),
		16, 0, 0, 0,
		currentDate.Location(),
	)

	for fridayDate.Weekday() != time.Friday {
		fridayDate = fridayDate.Add(time.Hour * HoursPerDay)
	}
	return fridayDate
}

// Write the additional code below to find if it's Saturday or Sunday
func isWeekend(currentDate time.Time) bool {
	return currentDate.Weekday() == time.Saturday || currentDate.Weekday() == time.Sunday
}

// Write the additional code below to find if it's Friday after 16:00
func isFridayAfternoon(currentDate time.Time) bool {
	return currentDate.Weekday() == time.Friday && currentDate.Hour() >= 16
}

// Finish the implementation of the main function below:
func main() {
	var currentDateInput string
	fmt.Scan(&currentDateInput)
	currentDate, err := time.Parse(time.RFC3339, currentDateInput)
	must(err)

	// Check if its Saturday or Sunday:
	if isWeekend(currentDate) {
		fmt.Println("Relax! It's a weekend day!")
		return
	}
	// Check if its Friday afternoon (after 16:00):
	if isFridayAfternoon(currentDate) {
		fmt.Println("Working week is over!")
		return
	}
	// Get the date of the next Friday:
	fridayDate := getFridayDate(currentDate)

	// Calculate the difference between the current date and the next Friday:
	difference := fridayDate.Sub(currentDate)
	minutes := int(difference.Minutes())

	hours := minutes / MinutesPerHour
	minutes %= MinutesPerHour

	fmt.Printf("It's %d hours and %d minutes until the next Friday!", hours, minutes)
}

// DO NOT delete the must() function — it is a helper to check for errors!
func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
