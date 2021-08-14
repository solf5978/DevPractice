package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, 11, 29, 21, 0, 0, 0, time.Local)
	fmt.Printf("Extract Units From: %v\n", t)

	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()

	fmt.Printf("The %dth day of %v is %v\n", dOfMonth, month, weekDay)
}
