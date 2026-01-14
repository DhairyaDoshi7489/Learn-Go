package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkUrl(url string, wg *sync.WaitGroup, results chan string) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		results <- fmt.Sprintf("[ERROR] %s is down!", url)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	// Send a formatted string into the channel
	results <- fmt.Sprintf("[OK] %s - %v", url, duration)
}

func main() {
	websites := []string{
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	var wg sync.WaitGroup

	results := make(chan string, len(websites))

	fmt.Println("Checking websites concurrently...")

	for _, url := range websites {
		wg.Add(1)
		go checkUrl(url, &wg, results)
	}
	wg.Wait()
	close(results)
	for msg := range results {
		fmt.Println(msg)
	}
}
