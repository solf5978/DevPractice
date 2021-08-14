package main

import (
	"fmt"
	"strconv"
)

const (
	bin         = "00001"
	hex         = "2f"
	intString   = "12"
	floatString = "12.3"
)

func main() {
	//	Decimals
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	//	Hexa Decimals
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed hexadecima: %d\n", res64)

	//	Binary
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed bin: %d\n", resBin)

	//	Floating-Points
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed float: %.5f\n", resFloat)
}
