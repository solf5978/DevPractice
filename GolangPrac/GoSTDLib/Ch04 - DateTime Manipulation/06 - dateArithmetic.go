package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		panic(err)
	}

	t := time.Date(2017, 11, 30, 1, 10, 20, 0, l)
	fmt.Printf("Default date is: %v\n", t)

	//	Add 3 days
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Default date + 3HRS is: %v\n", r1)

	//	Subtract 3 days
	r2 := t.Add(-72 * time.Hour)
	fmt.Printf("Defautl date - 3HRS is: %v\n", r2)

	//	w/ API Adding More Time
	r3 := t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1 YR + 3MTH + 2D is: %v\n", r3)
}
