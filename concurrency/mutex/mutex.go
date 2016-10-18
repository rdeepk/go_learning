//This packege is just for learning and experimenting with mutex.
//It runs various solutions for concurrent memory access.
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)
	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend)
				logtime := time.Now().Format(time.RFC3339)
				f.WriteString(logtime + " " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			go func(i, j int) {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}
	fmt.Scanln()
}
