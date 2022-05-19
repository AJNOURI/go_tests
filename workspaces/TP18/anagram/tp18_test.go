package anagram

import "testing"

func TestSimpleAnagram(t *testing.T) {
	// Given
	s1,s2 := "silent", "listen"
	expected := true
	
	// When
	result := areAnagrams(s1,s2)
	
	// Then
	if result != expected {
		t.Fail()
	}
}

func TestAnagramWithSpaces(t *testing.T) {
	// Given
	s1,s2 := "william Shakespeare", "i am a weakish speller"
	expected := true
	
	// When
	result := areAnagrams(s1,s2)
	
	// Then
	if result != expected {
		t.Fail()
	}
}


func TestBadAnagram(t *testing.T) {
	// Given
	s1,s2 := "this is not", "an anagram"
	expected := false
	
	// When
	result := areAnagrams(s1,s2)
	
	// Then
	if result != expected {
		t.Fail()
	}
}

