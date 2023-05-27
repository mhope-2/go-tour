package main

import (
	"fmt"
	"math/cmplx"
)

var is_protected bool

// factored vars
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const val = 89

func main(){
	var value int

	var i, j = 1, 2

	// := only available in funcs

	fmt.Println(value, is_protected)
	fmt.Println(i, j)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("const val", val)
}