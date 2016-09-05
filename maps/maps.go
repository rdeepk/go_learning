package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	str := make(map[string]int)
	sl := strings.Fields(s)
	for _, val := range sl {
		str[val]++
	}
	return str
}

func main() {
	wc.Test(WordCount)
}
