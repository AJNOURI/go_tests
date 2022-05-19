package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for n := range numberGenerator(context.TODO()) {
		fmt.Println(n)
		if n == 8 {
			break
		}
	}
}

func numberGenerator(ctx context.Context) <-chan int {
	dst := make(chan int)
	random := rand.New(rand.NewSource(time.Now().Unix()))
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dst)
				return
			case dst <- random.Intn(10):
				time.Sleep(time.Second)
			}
		}
	}()
	return dst
}
