package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Iteration", i, ":", z)
	}

	return z
}

func main() {
	vals := [5]float64{9, 10, 100, 300, 1000}

	for _, val := range vals {
		fmt.Println("==============")

		res := sqrt(val)
		fmt.Println("Square root of", val, ":", res, " VS Actual: ", math.Sqrt(val))
	}
}
