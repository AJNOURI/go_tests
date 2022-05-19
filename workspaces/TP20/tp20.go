package main

import (
	"fmt"
	"formation-go/TP20/anagram"
)

func main() {
	fmt.Println(anagram.AreAnagrams("this is not", "an anagram"))
	fmt.Println(anagram.AreAnagrams("mary", "army"))
	fmt.Println(anagram.AreAnagrams("silent", "listen"))
	fmt.Println(anagram.AreAnagrams("Madam Curie", "Radium came"))
	fmt.Println(anagram.AreAnagrams("william Shakespeare", "I am a weakish speller"))
}
