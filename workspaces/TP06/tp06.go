package main

import (
	"fmt"
)

// Define 2 new types Euros and Dollars, based on float64

const eurosDollarsRate = 1.168907

var (
	myEuros    Euros   = 123.45
	myDollars  Dollars = 678.9
	myBitcoins         = 0.2
)

func main() {
	fmt.Printf("I have %s, %s and %s\n", formatMoney(myEuros), formatMoney(myDollars), formatMoney(myBitcoins))
	fmt.Printf("My %s amount to %s\n", formatMoney(myEuros), formatMoney(eurosToDollars(myEuros)))
}

func formatMoney(val interface{}) string {
	// Use a type switch to format money values with the correct sign
	return ""
}

func eurosToDollars(euros Euros) Dollars {
	// Use type conversion to return the converted amount
	return 0
}
