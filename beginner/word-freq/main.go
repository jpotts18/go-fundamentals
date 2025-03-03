package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)


func cleanSentence (in string) string {
	lowerSentence := strings.ToLower(in)
	reg, err := regexp.Compile("[^a-z ]+")
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return ""
	}
	cleanedSentence := reg.ReplaceAllString(lowerSentence,"")
	return cleanedSentence
}

func main () {
	store := make(map[string]int)
	sentence := "The quick brown fox jumps over the lazy dog. The dog barks, but the fox keeps running.";
	cleaned := cleanSentence(sentence)
	words := strings.Split(cleaned, " ")
	for _, word := range words {
		if word != "" {
		store[word]++
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
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Key < pairs[j].Key
		}
		return pairs[i].Value > pairs[j].Value 
	})

	fmt.Println("Results")
	fmt.Println("-----------")
	for _, kv := range pairs {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
