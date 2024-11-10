package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl function recursively fetches URLs, ensuring URLs are fetched in parallel and only once
func Crawl(url string, depth int, fetcher Fetcher, visited map[string]bool, mu *sync.Mutex, wg *sync.WaitGroup) {
	// Ensure the current URL is not fetched again by locking the visited map
	mu.Lock()
	if visited[url] {
		mu.Unlock()
		return
	}
	visited[url] = true
	mu.Unlock()

	// Base case: If depth is 0, return without fetching more URLs
	if depth <= 0 {
		return
	}

	// Fetch the URL
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the URL and its body
	fmt.Printf("found: %s %q\n", url, body)

	// Use goroutines to fetch the URLs in parallel
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher, visited, mu, wg)
		}(u)
	}
}

func main() {
	// Initialize a mutex and wait group for concurrent access and waiting
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Create a visited map to keep track of visited URLs
	visited := make(map[string]bool)

	// Start the crawl from the root URL with depth 4
	Crawl("https://golang.org/", 4, fetcher, visited, &mu, &wg)

	// Wait for all goroutines to complete
	wg.Wait()
}

// fakeFetcher simulates a real fetcher for testing purposes
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// Fetch method simulates fetching a URL and returns its body and URLs
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fakeFetcher instance with predefined URLs for testing
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
