package main

import (
	"github.com/tour/wc"
	"strings"
)

// https://golang.org/pkg/strings/#Fields

func WordCount(s string) map[string]int {

	var words = make(map[string]int)
	tokens := strings.Split(s, " ")

	for _, v := range tokens {
		key := string(v)
		value := words[key]
		words[key] = value + 1
	}

	return words
}

func main() {
	wc.Test(WordCount)
}
