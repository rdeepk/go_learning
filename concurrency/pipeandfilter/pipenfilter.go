//This package generates prime numbers using goroutines and channels
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	ch := make(chan int)
	go generate(ch)
	for {
		ch1 := make(chan int)
		prime := <-ch
		fmt.Println(prime)
		go filter(ch, ch1, prime)
		ch = ch1	
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
		//fmt.Println("Generate Channel: ", i)
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		//fmt.Printf("In Filter: in = %d prime = %d\n", i, prime)
		if i%prime != 0 {
			//fmt.Println("Filter output is ", i)
			out <- i
		}
	}
}
