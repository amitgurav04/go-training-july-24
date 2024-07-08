package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.sdbkewbb.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for ch := range c {
		fmt.Println(ch)
	}

	for range links {
		fmt.Println(<-c)
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		ch <- link + " is down"
		fmt.Println(<-ch)
		return
	}
	ch <- link + " is up"
}
