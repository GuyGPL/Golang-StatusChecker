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
		"http://golang.org",
		"http://amazon.com",
	}

	// create channel (communicate between go routine)
	c := make(chan string)

	for _, link := range links {
		// create go routine
		go checkLink(link, c)
	}

	for l := range c {
		// if not pass 'l' in function literal 'l' can be difference in checklist() because 'l' refference to ram addess of main routine
		// ex. l inline 38 can be 'google.com' but in 42 can be 'facebook.com' 
		go func(l string) {
			time.Sleep(time.Second)
			checkLink(l, c)
		} (l)
		}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
