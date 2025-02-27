package main

import (
	"fmt"
	"regexp"
	"strings"
)

func cleanSentence (in string) string {
	lower_sentence := strings.ToLower(in)
	reg, _ := regexp.Compile("[^a-z ]+")
	cleaned_sentence := reg.ReplaceAllString(lower_sentence,"")
	return cleaned_sentence
}

func main () {
	store := make(map[string]int)
	sentence := "The quick brow fox jumps over the lazy dog. The dog barks, but the fox keeps running.";
	cleaned := cleanSentence(sentence)
	words := strings.Split(cleaned, " ")
	for _, word := range words {
		_, ok := store[word]
		if ok {
			store[word] = store[word] + 1
		} else {
			store[word] = 1
		}
	}

	fmt.Println(store)
}
