package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	link := "https://www.google.com"
	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}
	//fmt.Println(resp)

	bs := make([]byte, 99999)

	n, _ := resp.Body.Read(bs)
	fmt.Println(string(bs[:n]))

	io.Copy(os.Stdout, resp.Body)
}
