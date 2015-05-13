package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, t := 1.0, 0.0
	for math.Abs(z-t) > 1e-9 {
		t = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
