package concurrency

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var wgMutex sync.WaitGroup //pointer
var mut sync.Mutex         // pointer
var signals = []string{"https://linkedin.com"}

func Mutex() {
	websites := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
	}

	for _, url := range websites {
		go statusCode(url)
		wgMutex.Add(1)
	}

	wgMutex.Wait()
	fmt.Println(signals)

}

// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.

// A Mutex must not be copied after first use.

// In the terminology of the Go memory model, the n'th call to Unlock “synchronizes before” the m'th call to Lock for any n < m. A successful call to TryLock is equivalent to a call to Lock. A failed call to TryLock does not establish any “synchronizes before” relation at all.

func statusCode(url string) {
	defer wgMutex.Done()
	res, err := http.Get(url)

	parts := strings.Split(url, "/")
	domain := parts[len(parts)-2] // -2 to skip the empty string after the last "/"

	if err != nil {
		fmt.Println("Status Not Found")
	} else {
		signals = append(signals, domain)
		mut.Lock()
		fmt.Printf("Status Code form %s is %d \n", domain,
			res.StatusCode)
		mut.Unlock()
	}

}
