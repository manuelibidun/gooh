package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	l := <-c
	// 	fmt.Println(l)
	// 	go checkLink(l, c)
	// }
	// for {
	// 	link := <-c
	// 	go checkLink(link, c)
	// }
	for link := range c {
		// go checkLink(link, c)
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
