package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func main()  {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int{
	words := strings.Split(s, " ")
	wordCounts := make(map[string]int)
	for _, val := range words{
		_ , ok := wordCounts[val]
		if ok {
			wordCounts[val] += 1
		} else {
			wordCounts[val] = 1
		}
	}
	return wordCounts
}