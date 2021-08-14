package main

import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {
	//	Common way how to print the decimal number
	fmt.Printf("%d \n", integer)

	//	Always show the sign
	fmt.Printf("%+d \n", integer)

	//	Print in other base X16 o8 b2 d10
	fmt.Printf("%X \n", integer)
	fmt.Printf("%#X \n", integer)

	//	Padding with leading zeros
	fmt.Printf("%010d \n", integer)

	//	Left padding with spaces
	fmt.Printf("% 10d \n", integer)
	//	Right padding with spaces
	fmt.Printf("% -10d \n", integer)

	//	Print floating point number
	fmt.Printf("%f \n", floatNum)

	//	Floating point number with limited precision = 5
	fmt.Printf("%.5f \n", floatNum)

	//	Floating point in scientific notation
	fmt.Printf("%e \n", floatNum)

	//	Floating point number %e for large exponents or %f otherwise
	fmt.Printf("%g \n", floatNum)
}
