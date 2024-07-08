package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.sdbkewbb.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
	}
	var wg sync.WaitGroup
	wg.Add(1)
	for _, link := range links {
		go checkLink(link, &wg)
	}
	wg.Wait()
}

func checkLink(link string, wg *sync.WaitGroup) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		wg.Done()
		return
	}
	fmt.Println(link, "is up")
	wg.Done()
}
