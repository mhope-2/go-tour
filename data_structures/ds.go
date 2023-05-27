package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Structs
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)

	// change vec values through a pointer
	ptr := &v
	ptr.X = 2
	fmt.Println(v)

	// arrays & slices
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[2:5]
	fmt.Printf("slice: %d\n", s)

	sl := []bool{true, false, true}
	fmt.Println(sl)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(board)

	// appending to a slice
	fmt.Println(append([]int{1, 2, 3, 4}, 8, 9, 0))

	ss := []float64{12.43, 12.34, 12.34}
	for i, v := range ss {
		fmt.Printf("index: %v  value: %v\n", i, v)
	}

	// maps
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{3, 4}
	fmt.Println(m["Bell Labs"])

	fmt.Println(map[string]Vertex{
		"Bell Labs": {3, 5},
		"Google":    {2, 8},
	})

	mp := make(map[string]int)

	mp["Answer"] = 42
	fmt.Println("The value:", mp["Answer"])

	mp["Answer"] = 48
	fmt.Println("The value:", mp["Answer"])

	delete(mp, "Answer")
	fmt.Println("The value:", mp["Answer"])

	// check if key is present
	val, ok := mp["Answer"]
	if !ok {
		fmt.Printf("key: %v is not present\n", val)
	}

}
