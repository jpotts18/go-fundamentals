package main

import (
	"fmt"
	"regexp"
	"sort"
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
	sentence := "The quick brown fox jumps over the lazy dog. The dog barks, but the fox keeps running.";
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
	type kv struct {
		Key string
		Value int
	}

	var pairs []kv
	for k,v :=range store {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs,func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value 
	})

	fmt.Println("Results")
	fmt.Println("-----------")
	for _, kv := range pairs {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
