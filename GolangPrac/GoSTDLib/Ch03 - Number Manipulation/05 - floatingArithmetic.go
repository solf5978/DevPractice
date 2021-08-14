package main

import (
	"fmt"
	"math/big"
)

const (
	PI        = `3.1415926893284239754258405937539654248923`
	diameter  = 3.0
	precision = 400
)

func main() {
	pi, _ := new(big.Float).SetPrec(precision).SetString(PI)
	d := new(big.Float).SetPrec(precision).SetFloat64(diameter)

	circumference := new(big.Float).Mul(pi, d)

	pi64, _ := pi.Float64()
	fmt.Printf("Circumference big.Float = %.40f\n", circumference)
	fmt.Printf("Circumference float64 = %.40f\n", pi64*diameter)

	sum := new(big.Float).Add(pi, pi)
	fmt.Printf("Sum = %.4f\n", sum)

	diff := new(big.Float).Sub(pi, pi)
	fmt.Printf("Diff = %.4f\n", diff)

	quo := new(big.Float).Quo(pi, pi)
	fmt.Printf("Quocient = %.4f\n", quo)
}
