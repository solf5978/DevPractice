package main

import (
	"fmt"
	"math"
)

const (
	da        = 0.2999999995489375482932332432
	db        = 0.3
	TOLERANCE = 1e-8
)

func main() {
	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("Strings %s = %s equals: %v \n", daStr, dbStr, daStr == dbStr)
	fmt.Printf("Number equals: %v \n", db == da)

	//	As the precision of float representation is limited.
	//	For the float comparison it is better to use comparison with some tolerance.
	fmt.Printf("Number equals with Tolerance: %v \n", equals(da, db))
}

//	Equals compares the floating-point numbers with tolerance 1e-8
func equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < TOLERANCE {
		fmt.Print(delta)
		return true
	}
	return false
}
