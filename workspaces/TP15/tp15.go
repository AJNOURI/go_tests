package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type url struct {
	name string
	url  string
}

var urls = [...]url{
	url{"zenika", "http://zenika.com"},
	url{"microsoft", "http://microsoft.com"},
	url{"apple", "http://apple.com"},
	url{"facebook", "http://facebook.com"},
	url{"amazon", "http://amazon.com"},
}

func curl(u url) {
	t := time.Now()
	r, err := http.Get(u.url)
	if err != nil {
		fmt.Printf("Error fetching \"%s\":\n%s\n", u.url, err)
		return
	}
	defer r.Body.Close()
	n, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		fmt.Printf("Error processing \"%s\" body:\n%s\n", u.url, err)
		return
	}
	fmt.Printf("%s %d [%.2fs]\n", u.name, n, time.Since(t).Seconds())
}

func main() {
	start := time.Now()
	for _, u := range urls {
		go curl(u)
	}
	time.Sleep(6 * time.Second)
	fmt.Printf("Time to fetch all urls: %.0fs\n", time.Since(start).Seconds())
}
