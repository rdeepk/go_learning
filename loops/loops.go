package main

import (
	"fmt"
	"math"
)

const DELTA = 0.0000001
const INITIAL = 1.0

func Sqrt(x float64) (z float64) {
	z = INITIAL
	i := 1
	step := func() float64 {
		fmt.Printf("Ran %v iterations\n", i)
		i++
		return z - ( ((z*z) - x) / (2*z) )
	}
	for zz:= step(); math.Abs(zz - z) > DELTA; { 
		z = zz
		zz = step()
	}
	return
}

func main() {
	fmt.Println(math.Sqrt(4))
	fmt.Println(Sqrt(4))
}
