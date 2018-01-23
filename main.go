package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://twitter.com",
		"https://github.com",
		"https://golang.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, c)
		}(link)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println("Offline", l)
		c <- l

		return
	}

	fmt.Println("Online", l)
	c <- l
}
