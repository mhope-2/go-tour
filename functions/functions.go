package main

import (
	"fmt"
	"math"
)

// can omit type of x since x & y share same type
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// named return values (returns x, y)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Println(add(12, -7))
	fmt.Println(add(8, 5))

	a, b := swap("World", "Hello")
	fmt.Println(a, b)

	fmt.Println(split(18))

	// function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
