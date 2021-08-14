package main

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {

	//	Bad assumption on rounding the number by casting it to integer.

	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v\n", intVal)

	fRound := Round(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)

	fmt.Printf("Round by Math.Round %v\n", math.Round(valA))
	fmt.Printf("Round by Math.RoundToEven %v\n", math.RoundToEven(valA))
}

//	Round returns the nearest integer.
func Round(inputNum float64) float64 {
	truncNum := math.Trunc(inputNum)
	if math.Abs(inputNum-truncNum) >= 0.5 {
		return truncNum + math.Copysign(1, inputNum)
	}
	return truncNum
}
