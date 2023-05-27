package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)


func main(){
	total := 0

	for i := 0; i <= 10; i++{
		total += i;
	}

	fmt.Printf("total: %v\n", total)


	// The init and post statements are optional.
	sum := 1
	for ; sum <= 10; {
		sum += sum;
	}

	fmt.Printf("sum: %v\n", sum)

	// infinite for loop
	// this loop will spin, using 100% CPU (SA5002) go-staticcheck
	// for {
	// 	fmt.Println("Zazu")
	// }

	is_checked := false

	if is_checked{
		fmt.Println("checked")
	}else{
		fmt.Println("not checked")
	}

	if a := math.Pow(2, 5); a <10 {
		fmt.Println("false")
	} else{
		fmt.Println("true")
	}

	// switch
	// Switch without a condition is the same as switch true.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Print("OS X\n")
		case "linux":
			fmt.Print("Linux\n")
		default:
			fmt.Printf("%s\n", os) 
	}

	fmt.Println(time.Saturday)

	// Defer
	defer fmt.Print(" World\n")
	fmt.Print("Hello")

	fmt.Println("counting")

	// deferred calls are executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}