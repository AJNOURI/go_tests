package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var count = 0

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func visit(path string, f os.FileInfo, err error) error {
	// println(path)
	count++
	return nil
}

func doWalk(root string, c chan<- error) {
	ok := filepath.Walk(root, visit)
	c <- ok
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	go spinner(100 * time.Millisecond)

	c := make(chan error)
	go doWalk(root, c)
	<-c

	fmt.Printf("\nfilepath.Walk() visited %d elements\n", count)
}
