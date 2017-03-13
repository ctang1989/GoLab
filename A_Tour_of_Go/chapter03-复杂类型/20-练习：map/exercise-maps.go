package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	// initialize the wordMap variable
	wordMap := make(map[string]int)

	// separate the strings into words
	words := strings.Fields(s)

	// iterate over the words to count each instance
	for _, word := range words {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}

/*
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"quick":1, "jumped":1, "over":1, "dog.":1, "The":1, "brown":1, "fox":1, "the":1, "lazy":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"ate":2, "a":1, "donut.":2, "Then":1, "another":1, "I":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"canal":1, "panama.":1, "A":1, "man":1, "a":2, "plan":1}
*/
