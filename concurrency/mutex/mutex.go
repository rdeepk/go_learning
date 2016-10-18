//This packege is just for learning and experimenting with mutex.
//It runs various solutions for concurrent memory access.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	mutex := new(sync.Mutex)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func() {
				//the access of i and j in goroutines is wrong but this use is //just to arise the case of cuncurrent memory access
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}
	fmt.Scanln()
}
