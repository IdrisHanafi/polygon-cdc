package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// example 1
func swap(x, y string) (string, string) {
	return y, x
}

func example1() {
	a, b := swap("hello", "world")

	fmt.Println(a, b)
}

// example 2
// var c, python, java bool
var c, python, java = true, false, "no!"

func example2() {
	var i int

	j := 1 // implicit and not available on pkg level

	fmt.Println(i, j, c, python, java)
}

// example 3
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func example3() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// example 4
func example4() {
	var x, y int = 3, 4

	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)

	fmt.Println(x, y, z)
}

// example 5
func loops1(numToPrint int) {
	for i := 0; i < numToPrint; i++ {
		fmt.Println("hi", i, "times")
	}
}

// example 6: a while loop
func loops2() {
	sums := 10
	for sums < 1000 {
		sums += 100
	}
	fmt.Println("sum:", sums)
}

// example 7
func sqrt(num float64) string {
	if num < 0 {
		return sqrt(-num) + "i"
	}

	return fmt.Sprint(math.Sqrt(num))
}

func example7() {
	fmt.Println("Square root of 1: ", sqrt(1))
	fmt.Println("Square root of 0: ", sqrt(0))
	fmt.Println("Square root of -10: ", sqrt(-10))
	fmt.Println("Square root of 10: ", sqrt(10))
}

func main() {
	example1()
	example2()
	example3()
	example4()
	loops1(5)
	loops2()
	example7()
}
