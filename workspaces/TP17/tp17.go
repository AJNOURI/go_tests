package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(areAnagrams("this is not", "an anagram"))
	fmt.Println(areAnagrams("mary", "army"))
	fmt.Println(areAnagrams("silent", "listen"))
	fmt.Println(areAnagrams("Madam Curie", "Radium came"))
	fmt.Println(areAnagrams("william Shakespeare", "I am a weakish speller"))
}

func areAnagrams(a, b string) bool {
	a = normalize(a)
	b = normalize(b)

	if len(a) != len(b) {
		return false
	}

	for _, c := range a {
		var i int
		if i = strings.IndexRune(b, c); i == -1 {
			return false
		}
		b = removeCharAtIndex(b, i)
	}
	return true
}

func normalize(s string) (result string) {
	result = removeWhiteSpaces(toLower(s))
	return
}

func toLower(s string) (result string) {
	return strings.ToLower(s)
}

func removeWhiteSpaces(s string) (result string) {
	return strings.Join(strings.Fields(s), "")
}

func removeCharAtIndex(s string, i int) (result string) {
	var a, b string

	a = s[:i]
	if i < len(s) {
		b = s[i+1:]
	}

	result = a + b
	return
}
