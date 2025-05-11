package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Split(s, " ")
	var word_count = make(map[string]int)
	for _, v := range words {
		elm, ok := word_count[v]
		if ok {
			word_count[v] = elm + 1

		} else {
			word_count[v] = 1
		}

	}
	return word_count
}

func main() {
	wc.Test(WordCount)
}
