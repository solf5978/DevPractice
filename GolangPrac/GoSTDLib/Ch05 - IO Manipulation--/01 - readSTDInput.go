package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("How old are you?")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hello %s yor are %d now\n", name, age)
}
