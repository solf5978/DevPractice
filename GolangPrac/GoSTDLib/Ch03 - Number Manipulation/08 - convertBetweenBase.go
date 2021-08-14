package main

import (
	"fmt"
	"strconv"
)

const (
	bin      = "10111"
	hex      = "1A"
	oct      = "12"
	dec      = "10"
	floatNum = 16.123557
)

func main() {

	t1, _ := strconv.ParseInt(oct, 8, 64)
	fmt.Printf("Test %s Output%s\n", oct, strconv.FormatInt(t1, 2))
	//	Converts binary value into hex
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("Binary value %s converted to hex: %s\n", bin, v)

	//	Converts hex value into dec
	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("Hex value %s converted to dec: %s\n", hex, v)

	//	Converts oct value into hex
	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("Oct value %s converted to hex: %s\n", oct, v)

	//	Converts dec value into oct
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Dec value %s converted to oct: %s \n", dec, v)
}

//	ConvertInt converts the given string value of base to defined toBase.
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
