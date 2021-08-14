package main

import (
	"fmt"
	"strings"
)

const refString = "Mary*had,a%little_lamb"

func main() {
	//	The splitFunc is called for each rune(untyped int32) in a string.
	//	If the rune equals any of character in a "*%,_"
	//	the refString is split.

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}
	words := strings.FieldsFunc(refString, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}
