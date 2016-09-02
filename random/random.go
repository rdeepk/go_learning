package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum/2
	y = sum-6
	return
}
func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)
}

