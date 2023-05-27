package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("my favorite number is", rand.Intn(10))
	fmt.Printf("square root of 7 is %g\n", math.Sqrt(7))
	fmt.Printf("pi val: %g", math.Pi)
}
