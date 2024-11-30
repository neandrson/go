package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {
	nextDay := start.AddDate(0, 0, 1)

	for nextDay.Weekday() == time.Saturday || nextDay.Weekday() == time.Sunday {
		nextDay = nextDay.AddDate(0, 0, 1)
	}

	return nextDay
}

func main() {
	startTime := time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC)
	nextWorkday := NextWorkday(startTime)

	fmt.Println(nextWorkday)
}
