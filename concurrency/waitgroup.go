package concurrency

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup // uncomment

func Waitgroup() {
	websites := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.anik.com/",
	}

	for _, url := range websites {
		//		getStatusCode(url) // comment
		wg.Add(1)             // uncomment
		go getStatusCode(url) // uncomment

	}

	wg.Wait() // uncomment

}

func getStatusCode(url string) {
	defer wg.Done() // uncomment
	res, err := http.Get(url)

	parts := strings.Split(url, "/")
	domain := parts[len(parts)-2] // -2 to skip the empty string after the last "/"

	if err != nil {
		fmt.Println("Status Not Found")
	} else {
		fmt.Printf("Status Code form %s is %d \n", domain, res.StatusCode)
	}

}
