package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func FetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response from %s: %v\n", url, err)
		return
	}
	fmt.Printf("Fetched %s in %v, size: %d bytes\n", url, time.Since(start), len(body))

}

func main() {
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://golang.org",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go FetchURL(url, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished execution")
}
