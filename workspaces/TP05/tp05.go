package main

import "fmt"

func count(from, to int) {
	i := from
	for {
		fmt.Println(i)
		if i == to {
			break
		}
		i++
	}
}

func main() {
	count(0, 10)
	fmt.Println("done!")
}
