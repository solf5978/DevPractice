package main

import (
	"fmt"
	"time"
)

func main() {
	//	If timezone is not defined than Parse function returns the time in UTC timezone.
	t, err := time.Parse("2/1/2006", time.Now().Local().Format("2/1/2006"))
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	//	If timezone is given than it is parsed in given timezone
	t, err = time.Parse("2/1/2006 3:04 PM MST", "31/7/2015 1:25 AM DST")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	//	Note that the ParseInLocation parses the time in give location,
	//	if the string does not contain time zone definition
	t, err = time.ParseInLocation("2/1/2006 3:04 PM ", "31/7/2015 1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
