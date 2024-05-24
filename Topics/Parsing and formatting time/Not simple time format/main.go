package main

import (
	"fmt"
	"time"
)

func main() {
	var year, month, day, hour, minute int

	fmt.Scan(&year, &month, &day, &hour, &minute)

	date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	fmt.Println(date.Format("-* 02 | 01 | 2006 *-\n   < 15 >< 04 >"))
}
