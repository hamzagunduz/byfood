package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var strArray = []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	countWords(strArray)
}

func countWords(words []string) {
	var mWords = make(map[string]int)
	var mWordsWithA = make(map[string]int)

	for _, word := range words {
		var counter int
		for _, letter := range word {
			if strings.Contains(string(letter), "a") {
				counter = counter + 1
			}
		}
		if counter > 0 {
			mWordsWithA[word] = counter
		} else {
			mWords[word] = len(word)
		}

	}

	sortWords(mWordsWithA)
	sortWords(mWords)
}
func sortWords(words map[string]int) {
	type strct struct {
		Key   string
		Value int
	}

	var sortSlice []strct
	for x, y := range words {
		sortSlice = append(sortSlice, strct{x, y})
	}

	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].Value > sortSlice[j].Value
	})

	for x, strct := range sortSlice {
		if x == 0 {
			fmt.Printf("%s", strct.Key)
		} else {
			fmt.Printf(", %s", strct.Key)
		}
	}
}
