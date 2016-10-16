//This package is just for experimenting with basic channels usage.
package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's soul\n"

	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}

	close(ch)

	for {
		if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	}
}
