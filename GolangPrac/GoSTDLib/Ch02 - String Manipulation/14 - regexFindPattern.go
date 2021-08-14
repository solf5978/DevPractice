package main

import (
	"fmt"
	"regexp"
)

const refString = `[{ \"email\": \"email@example.com\" \
"phone\": 555467890},
{ \"email\": \"other@domain.com\" \
"phone\": 555467890}]`

func main() {

	//	This pattern is simplified for brevity
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")

	first := emailRegexp.FindString(refString)
	fmt.Println("First: " + first)

	all := emailRegexp.FindAllString(refString, -1)
	fmt.Println("All: ", all)
	for _, val := range all {
		fmt.Println(val)
	}
}
