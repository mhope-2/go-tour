package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	p := &Vertex{3, 4}
	p.Scale(10)
	fmt.Println("p = ", p)
}
