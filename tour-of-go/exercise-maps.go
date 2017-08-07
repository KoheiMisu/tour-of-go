package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strings := strings.Fields(s)
	wordCountSet := make(map[string]int, len(strings))

	for _, str := range strings {
		if _, ok := wordCountSet[str]; ok {
			wordCountSet[str]++
		} else {
			wordCountSet[str] = 1
		}
	}

	return wordCountSet
}

func main() {
	wc.Test(WordCount)
}
