package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	strings := strings.Fields(s)
	for _,word := range strings {
		count, exist := wordMap[word]
		if exist {
			wordMap[word] = count + 1;
		} else {
			wordMap[word] = 1;
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
