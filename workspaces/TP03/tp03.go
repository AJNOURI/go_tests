package main

import "fmt"

var pi = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 9}

func isPi(digits ...int) bool {
	for i, digit := range digits {
		if pi[i] != digit {
			return false
		}
	}
	return true
}

func main() {

	if isPi(3, 1, 4) {
		fmt.Println("Yeah! These digits correspond to Pi beginning.")
	} else {
		fmt.Println("Nope! These digits do not correspond to Pi beginning.")
	}

	successMessage := func() { fmt.Println("Yeah! These digits correspond to Pi beginning.") }
	failureMessage := func() { fmt.Println("Nope! These digits do not correspond to Pi beginning.") }
	if isPi(3, 1, 4, 1) {
		successMessage()
	} else {
		failureMessage()
	}
}
