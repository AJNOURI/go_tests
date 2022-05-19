package main

import "fmt"

func count(from, to int) {
	i := from
	for {
		fmt.Println(i)
		// Each time it arrives to defer, it is pushed to called in FIFO
		// and ONLY when the container function count finished
		defer fmt.Println(i)
		if i == to {
			panic("panic")
			//break
		}
		i++
	}
}

func main() {
	count(0, 10)
	fmt.Println("done!")
}
