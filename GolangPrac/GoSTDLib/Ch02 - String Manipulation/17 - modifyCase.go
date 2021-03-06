package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	email = "ExamPle@domain.com"
	name  = "issac newton"
	upc   = "upc"
	i     = "i"

	snakeCase = "first_name"
)

func main() {

	//	For comparing the use input sometimes it is better to
	//	compare the input in a same case.

	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Println("Upper case:" + upcCode)

	//	This digraph has different uppder case and title case.
	str := "dzadasdf"
	fmt.Printf("%s in upper: %s and title: %s \n", str,
		strings.ToUpper(str), strings.ToTitle(str))

	//	Use of XXXSpecial function
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is different: %#U vs. %#U \n", title[0], []rune(titleTurk)[0])
	}

	//	In some cases the input needs to be corrected in case.
	correctNameCase := strings.Title(name)
	fmt.Println("Corrected name: ", correctNameCase)

	//	Converting the snake case to camel case with use of Title and ToLower functions.
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: ", firstNameCamel)
}

func toCamelCase(input string) string {
	titleSpeace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpeace, " ", "", -1)
	return strings.ToLower(camel[:1]) + camel[1:]
}
