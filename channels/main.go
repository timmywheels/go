package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string {
		"http://timwheeler.com",
		"http://sendpoint.io",
		"http://codesnippet.io",
		"http://scooped.io",
		"http://shop.scoopedd.io",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkStatus(url, c)
	}

	//for i := 0; i < len(urls); i++ {
	//	go checkStatus(<- c)
	//}
	for url := range c {
		//time.Sleep(5 * time.Second)
		//go checkStatus(url, c)
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkStatus(url, c)
		}(url)

	}
}

func checkStatus(url string, channel chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is down!")
		//channel <- url + " is down!"
		channel <- url
		return
	}
	channel <- url
	//channel <- url + " is up!"
	fmt.Println(url, "is up!")
}